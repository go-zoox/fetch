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
	"net/textproto"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/headers"
	"github.com/tidwall/gjson"

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
		if err := fmt.PrintJSON("[GOZOOX_FETCH][DEBUG][Request]", config); err != nil {
			fmt.Println("[warn] failed to fmt.PrintJSON:", err, config)
		}
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
		default:
			return nil, fmt.Errorf("unsupport proxy(%s)", config.Proxy)
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
	if config.Username != "" || config.Password != "" {
		req.URL.User = url.UserPassword(config.Username, config.Password)
	}

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

					if f, ok := v.(io.ReadCloser); ok {
						// fix multipart form file content type
						if err := createFormFile(w, f, k); err != nil {
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
		} else if strings.Contains(req.Header.Get(headers.ContentType), "application/octet-stream") {
			if config.Body == nil {
				return nil, fmt.Errorf("octet-stream body is required")
			}

			body, ok := config.Body.(io.ReadCloser)
			if !ok && body != nil {
				body = io.NopCloser(body)
			}

			req.Body = body
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

	if config.CompressRequest {
		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)

		// Compress the request body
		if _, err := io.Copy(gz, req.Body); err != nil {
			return nil, fmt.Errorf("failed to compress request body: %v", err)
		}
		if err := gz.Close(); err != nil {
			return nil, fmt.Errorf("failed to close gzip writer: %v", err)
		}

		req.Header.Set(headers.ContentEncoding, "gzip")
		req.Body = io.NopCloser(&buf)
		req.Header.Set(headers.ContentLength, strconv.Itoa(buf.Len()))
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

	if config.BasicAuth.Username != "" || config.BasicAuth.Password != "" {
		f.SetBasicAuth(config.BasicAuth.Username, config.BasicAuth.Password)
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
				Total:    resp.ContentLength,
				Current:  0,
				Reporter: f.config.OnProgress,
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

	if os.Getenv(EnvDEBUG) != "" {
		if strings.Contains(resp.Header.Get(headers.ContentType), "application/json") {
			b, err := json.MarshalIndent(gjson.Parse(string(body)).Value(), "", "  ")
			if err != nil {
				fmt.Println("[GOZOOX_FETCH][DEBUG][Response]", string(body))
			} else {
				fmt.Println("[GOZOOX_FETCH][DEBUG][Response]", string(b))
			}
		} else {
			fmt.Println("[GOZOOX_FETCH][DEBUG][Response]", string(body))
		}
	}

	return &Response{
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Body:    body,
		//
		Request: config,
	}, nil
}

// @TODO for multipart/form-data with file
//
//		Issue:
//			mime/multipart:CreateFormFile has a fixed content type(application/octet-stream),
//			does not support auto-detect content type
//
//		Need: Create MIME encoded form files that auto-detect the content type.
//
//		Reference:
//	 	1. https://groups.google.com/g/golang-nuts/c/HwOYproYQqA
//	 	2. https://github.com/go-openapi/runtime/pull/170/files

// NamedReadCloser is a named reader
type NamedReadCloser interface {
	io.ReadCloser

	Name() string
}

func escapeQuotes(s string) string {
	return strings.NewReplacer("\\", "\\\\", `"`, "\\\"").Replace(s)
}

func createFormFile(w *multipart.Writer, reader io.ReadCloser, fieldname string) error {
	buf := bytes.NewBuffer([]byte{})
	filename := ""
	if f, ok := reader.(NamedReadCloser); ok {
		filename = f.Name()
	}

	// Need to read the data so that we can detect the content type
	if _, err := io.Copy(buf, reader); err != nil {
		return err
	}
	fileBytes := buf.Bytes()
	fileContentType := http.DetectContentType(fileBytes)

	newFi := CreateNamedReader(filename, buf)

	h := make(textproto.MIMEHeader)
	if filename == "" {
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldname)))
	} else {
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escapeQuotes(fieldname), escapeQuotes(filename)))
	}
	h.Set("Content-Type", fileContentType)

	fw, err := w.CreatePart(h)
	if err != nil {
		return err
	}

	if _, err = io.Copy(fw, newFi); err != nil {
		return err
	}

	return nil
}

// CreateNamedReader creates a named reader
//
//		multipart.File, that is Request.ParseMultipartForm, does not have a name
//		so we need to create a named reader to get the file name
//	 when uploading a file with multipart/form-data
func CreateNamedReader(name string, rdr io.Reader) NamedReadCloser {
	rc, ok := rdr.(io.ReadCloser)
	if !ok {
		rc = io.NopCloser(rdr)
	}
	return &namedReadCloser{
		name: name,
		cr:   rc,
	}
}

type namedReadCloser struct {
	name string
	cr   io.ReadCloser
}

func (n *namedReadCloser) Close() error {
	return n.cr.Close()
}
func (n *namedReadCloser) Read(p []byte) (int, error) {
	return n.cr.Read(p)
}
func (n *namedReadCloser) Name() string {
	return n.name
}
