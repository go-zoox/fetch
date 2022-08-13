package fetch

import (
	"bytes"
	"compress/gzip"
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

	client := &http.Client{
		Timeout: config.Timeout,
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
				TLSHandshakeTimeout: 10 * time.Second,
			}
		case "socks5", "socks5h":
			dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
			if err != nil {
				return nil, fmt.Errorf("invalid socks5 proxy: %s", config.Proxy)
			}

			client.Transport = &http.Transport{
				Proxy:               http.ProxyFromEnvironment,
				Dial:                dialer.Dial,
				TLSHandshakeTimeout: 10 * time.Second,
			}
		}
	}

	req, err := http.NewRequest(methodOrigin, fullURL, nil)
	if err != nil {
		// panic("error creating request: " + err.Error())
		return nil, errors.New(ErrCannotCreateRequest.Error() + ": " + err.Error())
	}

	// @TODO
	if _, ok := config.Body.(string); ok {
		req.Header.Set(HeaderContentTye, "text/plain")
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
		if req.Header.Get(HeaderContentTye) == "" {
			req.Header.Set(HeaderContentTye, "application/json")
		}

		if strings.Contains(req.Header.Get(HeaderContentTye), "application/json") {
			body, err := json.Marshal(config.Body)
			if err != nil {
				// panic("error marshalling body: " + err.Error())
				return nil, errors.New(ErrInvalidJSONBody.Error() + ": " + err.Error())
			}

			// req.Header.Set(HeaderContentTye, "application/json")
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		} else if strings.Contains(req.Header.Get(HeaderContentTye), "application/x-www-form-urlencoded") {
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
		} else if strings.Contains(req.Header.Get(HeaderContentTye), "multipart/form-data") {
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
				req.Header.Set(HeaderContentTye, w.FormDataContentType())
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
				req.Header.Set(HeaderContentTye, w.FormDataContentType())
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

	resp, err := client.Do(req)

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get(HeaderContentEncoding) {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("gzip decode error: %s", err)
		}
		// defer reader.Close()
	default:
		reader = resp.Body
	}

	if err != nil {
		// panic("error sending request: " + err.Error())
		return nil, errors.New(ErrSendingRequest.Error() + ": " + err.Error())
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

		_, err = io.Copy(file, reader)
		if err != nil {
			return nil, err
		}

		return &Response{
			Status:  resp.StatusCode,
			Headers: resp.Header,
			//
			Request: config,
		}, nil
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
		return nil, errors.New(ErrReadingResponse.Error() + ": " + err.Error())
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
