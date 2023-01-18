package fetch

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/headers"

	"golang.org/x/net/proxy"
)

// Execute executes the request
func (f *Fetch) Execute() (*Response, error) {
	if len(f.Errors) > 0 {
		return nil, f.Errors[0]
	}

	config, err := f.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %v", err)
	}

	if os.Getenv(EnvDEBUG) != "" {
		fmt.PrintJSON("[GOZOOX_FETCH][DEBUG][Request]", config)
	}

	fullURL := config.URL
	methodOrigin := config.Method
	// @ORIGIN QUERY
	var urlQueryOrigin url.Values
	if strings.ContainsAny(fullURL, "?") {
		u, err := url.Parse(config.BaseURL)
		if err != nil {
			return nil, errors.New("failed to parsed origin url")
		}

		urlQueryOrigin = u.Query()
	}

	if config.TLSCaCertFile != "" {
		caCrt, err := ioutil.ReadFile(config.TLSCaCertFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read tls certificate file(%s): %v", config.TLSCaCertFile, err)
		}

		config.TLSCaCert = caCrt
	}

	if config.TLSCertFile != "" {
		clientCrt, err := ioutil.ReadFile(config.TLSCertFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read tls certificate file(%s): %v", config.TLSCertFile, err)
		}

		config.TLSCert = clientCrt
	}

	if config.TLSKeyFile != "" {
		clientKey, err := ioutil.ReadFile(config.TLSKeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read tls certificate file(%s): %v", config.TLSKeyFile, err)
		}

		config.TLSKey = clientKey
	}

	transport := http.DefaultTransport
	if config.TLSCaCert != nil {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(config.TLSCaCert)

		// defaultTransportDialContext := func(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
		// 	return dialer.DialContext
		// }

		// transport = &http.Transport{
		// 	Proxy: http.ProxyFromEnvironment,
		// 	DialContext: defaultTransportDialContext(&net.Dialer{
		// 		Timeout:   30 * time.Second,
		// 		KeepAlive: 30 * time.Second,
		// 	}),
		// 	ForceAttemptHTTP2:     true,
		// 	MaxIdleConns:          100,
		// 	IdleConnTimeout:       90 * time.Second,
		// 	TLSHandshakeTimeout:   10 * time.Second,
		// 	ExpectContinueTimeout: 1 * time.Second,
		// 	// https://stackoverflow.com/questions/38822764/how-to-send-a-https-request-with-a-certificate-golang
		// 	TLSClientConfig: &tls.Config{
		// 		RootCAs: pool,
		// 	},
		// }

		tr := transport.(*http.Transport)
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}
		tr.TLSClientConfig.RootCAs = pool
	}

	if config.TLSCert != nil && config.TLSKey != nil {
		tr := transport.(*http.Transport)
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}

		clientCrt, err := tls.X509KeyPair(config.TLSCert, config.TLSKey)
		if err != nil {
			return nil, fmt.Errorf("failed to load client cert and key: %v", err)
		}

		tr.TLSClientConfig.Certificates = []tls.Certificate{clientCrt}
	}

	if config.TLSInsecureSkipVerify {
		tr := transport.(*http.Transport)
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}
		tr.TLSClientConfig.InsecureSkipVerify = config.TLSInsecureSkipVerify
	}

	// if f.config.HTTP2 {
	// 	if err := http2.ConfigureTransport(&transport); err != nil {
	// 		return nil, fmt.Errorf("failed to configure http2: %v", err)
	// 	}
	// }

	client := &http.Client{
		Timeout:   config.Timeout,
		Transport: transport,
	}

	// apply proxy
	if config.Proxy != "" {
		// fmt.Println("proxy:", config.Proxy)
		proxyURL, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy: %s", config.Proxy)
		}

		switch proxyURL.Scheme {
		case "http", "https":
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				// default transport
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}
		case "socks5", "socks5h":
			dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
			if err != nil {
				return nil, fmt.Errorf("invalid socks5 proxy: %s", config.Proxy)
			}

			client.Transport = &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial:  dialer.Dial,
				// default transport
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}
		}
	}

	req, err := http.NewRequestWithContext(f.config.Context, methodOrigin, fullURL, nil)
	if err != nil {
		// panic("error creating request: " + err.Error())
		return nil, errors.New("ErrCannotCreateRequest(1): " + ErrCannotCreateRequest.Error() + ", err: " + err.Error())
	}

	// @TODO
	if _, ok := config.Body.(string); ok {
		req.Header.Set(headers.ContentType, "text/plain")
	}

	for k, v := range config.Headers {
		// ignore empty value
		if v != "" {
			req.Header.Set(k, v)
		}
	}

	query := req.URL.Query()
	// apply origin query
	for k, v := range urlQueryOrigin {
		query.Add(k, v[0])
	}
	// apply custom query
	for k, v := range config.Query {
		// ignore empty value
		if v != "" {
			query.Add(k, v)
		}
	}
	req.URL.RawQuery = query.Encode()

	// if GET, ignore Body
	if config.Body != nil && config.Method == GET {
		// // panic("Cannot set body for GET request")
		// return nil, ErrCannotSendBodyWithGet
		config.Body = nil
	}

	if config.Body != nil {
		if req.Header.Get(headers.ContentType) == "" {
			req.Header.Set(headers.ContentType, "application/json")
		}

		if strings.Contains(req.Header.Get(headers.ContentType), "application/json") {
			body, err := json.Marshal(config.Body)
			if err != nil {
				// panic("error marshalling body: " + err.Error())
				return nil, errors.New("ErrInvalidJSONBody(2): " + ErrInvalidJSONBody.Error() + ", err: " + err.Error())
			}

			// req.Header.Set(HeaderContentTye, "application/json")
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		} else if strings.Contains(req.Header.Get(headers.ContentType), "application/x-www-form-urlencoded") {
			body := url.Values{}
			if kv, ok := config.Body.(map[string]string); ok {
				for k, v := range kv {
					body.Add(k, v)
				}
			} else {
				return nil, errors.New(ErrInvalidURLFormEncodedBody.Error() + ": must be map[string]string")
			}

			// req.Header.Set(HeaderContentTye, "application/x-www-form-urlencoded")
			// req.Body = ioutil.NopCloser(bytes.NewReader(body))
			req.Body = ioutil.NopCloser(strings.NewReader(body.Encode()))
		} else if strings.Contains(req.Header.Get(headers.ContentType), "multipart/form-data") {
			if values, ok := config.Body.(map[string]interface{}); ok {
				var b bytes.Buffer
				w := multipart.NewWriter(&b)
				for k, v := range values {
					if v == nil {
						continue
					}

					var fw io.Writer
					if text, ok := v.(string); ok {
						if fw, err = w.CreateFormField(k); err != nil {
							return nil, err
						}

						if _, err = io.Copy(fw, strings.NewReader(text)); err != nil {
							return nil, err
						}

						continue
					}

					if file, ok := v.(io.Reader); ok {
						// @TODO
						// filename := file.Name()
						filename := k
						type File interface {
							Name() string
						}

						if f, ok := file.(File); ok {
							filename = f.Name()
						}

						if fw, err = w.CreateFormFile(k, filename); err != nil {
							return nil, err
						}

						if _, err = io.Copy(fw, file); err != nil {
							return nil, err
						}
						continue
					}
				}
				w.Close()
				req.Header.Set(headers.ContentType, w.FormDataContentType())
				req.Body = ioutil.NopCloser(&b)
			} else if values, ok := config.Body.(map[string]string); ok {
				var b bytes.Buffer
				w := multipart.NewWriter(&b)
				for k, v := range values {
					var fw io.Writer
					if fw, err = w.CreateFormField(k); err != nil {
						return nil, err
					}

					if _, err = io.Copy(fw, strings.NewReader(v)); err != nil {
						return nil, err
					}

					continue
				}
				w.Close()
				req.Header.Set(headers.ContentType, w.FormDataContentType())
				req.Body = ioutil.NopCloser(&b)
			} else {
				return nil, errors.New(ErrInvalidBodyMultipart.Error() + ": must be map[string]interface{} or map[string]string")
			}
		} else {
			if _, ok := config.Body.(string); !ok {
				return nil, ErrorInvalidBody
			}

			req.Body = ioutil.NopCloser(bytes.NewReader([]byte(config.Body.(string))))
		}
	}

	// unix domain socket: https://gist.github.com/teknoraver/5ffacb8757330715bcbcc90e6d46ac74
	if config.UnixDomainSocket != "" {
		// remove unix://
		// if strings.HasPrefix(config.UnixDomainSocket, "unix://") {
		// 	config.UnixDomainSocket = config.UnixDomainSocket[7:]
		// }
		config.UnixDomainSocket = strings.TrimPrefix(config.UnixDomainSocket, "unix://")

		tr := client.Transport.(*http.Transport)
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial("unix", config.UnixDomainSocket)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		// panic("error sending request: " + err.Error())
		return nil, errors.New("ErrSendingRequest(3):  " + ErrSendingRequest.Error() + ", err: " + err.Error() + "(Please check your network, maybe use bad proxy or network offline)")
	}

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get(headers.ContentEncoding) {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("gzip decode error: %s", err)
		}
		// defer reader.Close()
	default:
		reader = resp.Body
	}

	if !config.IsStream {
		defer reader.Close()
	}

	if config.IsSession {
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			f.SetCookie(cookie.Name, cookie.Value)
		}
	}

	if config.DownloadFilePath != "" {
		file, err := os.OpenFile(config.DownloadFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		res := &Response{
			Status:  resp.StatusCode,
			Headers: resp.Header,
			//
			Request: config,
		}

		if f.config.OnProgress != nil {
			progress := &Progress{
				Total:   resp.ContentLength,
				Current: 0,
			}

			_, err = io.Copy(io.MultiWriter(file, progress), reader)
			if err != nil {
				return nil, err
			}
		} else {
			_, err = io.Copy(file, reader)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	}

	if config.IsStream {
		return &Response{
			Status:  resp.StatusCode,
			Headers: resp.Header,
			//
			Request: config,
			//
			Stream: reader,
		}, nil
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		// panic("error reading response: " + err.Error())
		return nil, errors.New("ErrReadingResponse(4): " + ErrReadingResponse.Error() + ", err: " + err.Error())
	}

	// fmt.Println("response: ", string(body))

	return &Response{
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Body:    body,
		//
		Request: config,
	}, nil
}
