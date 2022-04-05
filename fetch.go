package fetch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type Fetch struct {
	config *Config
	Errors []error
}

func New(cfg ...*Config) *Fetch {
	config := DefaultConfig()
	if len(cfg) > 1 {
		panic("Too many arguments")
	}

	if len(cfg) == 1 {
		config.Merge(cfg[0])
	}

	return &Fetch{config: config}
}

func (f *Fetch) SetConfig(configs ...*Config) *Fetch {
	for _, config := range configs {
		f.config.Merge(config)
	}

	return f
}

// @TODO
func (f *Fetch) getMethodConfig(config ...*Config) *Config {
	if len(config) > 0 {
		return config[0]
	}

	return &Config{}
}

func (f *Fetch) SetUrl(url string) *Fetch {
	f.config.Url = url
	return f
}

func (f *Fetch) SetMethod(method string) *Fetch {
	for m := range METHODS {
		if method == METHODS[m] {
			f.config.Method = method
			return f
		}
	}

	f.Errors = append(f.Errors, ErrInvalidMethod)
	return f
}

func (f *Fetch) SetHeader(key, value string) *Fetch {
	f.config.Headers[key] = value
	return f
}

func (f *Fetch) SetQuery(key, value string) *Fetch {
	f.config.Query[key] = value
	return f
}

func (f *Fetch) SetParam(key, value string) *Fetch {
	f.config.Params[key] = value
	return f
}

func (f *Fetch) SetBody(body ConfigBody) *Fetch {
	f.config.Body = body
	return f
}

//
func (f *Fetch) SetBaseURL(url string) *Fetch {
	f.config.BaseURL = url
	return f
}

func (f *Fetch) SetTimeout(timeout time.Duration) *Fetch {
	f.config.Timeout = timeout
	return f
}

func (f *Fetch) SetUserAgent(userAgent string) *Fetch {
	f.SetHeader("user-agent", userAgent)
	return f
}

func (f *Fetch) SetBasicAuth(username, password string) *Fetch {
	f.SetAuthorization("Basic " + base64.RawStdEncoding.EncodeToString([]byte(username+":"+password)))
	return f
}

func (f *Fetch) SetBearToken(token string) *Fetch {
	f.SetAuthorization("Bearer " + token)
	return f
}

func (f *Fetch) SetAuthorization(token string) *Fetch {
	f.SetHeader("Authorization", token)
	return f
}

//
func (f *Fetch) Execute() (*Response, error) {
	if len(f.Errors) > 0 {
		return nil, f.Errors[0]
	}

	methodOrigin := f.config.Method
	fullUrl := f.config.Url
	// @ORIGIN QUERY
	var urlQueryOrigin url.Values
	if strings.ContainsAny(fullUrl, "?") {
		u, err := url.Parse(f.config.BaseURL)
		if err != nil {
			return nil, errors.New("failed to parsed origin url")
		}

		fmt.Println("RawQuery:", u.RawQuery, u.RawQuery != "")
		urlQueryOrigin = u.Query()
	}

	// @BASEURL
	if f.config.BaseURL != "" {
		parsedBaseUrl, err := url.Parse(f.config.BaseURL)
		if err != nil {
			return nil, errors.New("invalid base URL")
		}

		parsedBaseUrl.Path = path.Join(parsedBaseUrl.Path, fullUrl)
		fullUrl = parsedBaseUrl.String()
	}

	client := &http.Client{
		Timeout: f.config.Timeout,
	}
	req, err := http.NewRequest(methodOrigin, fullUrl, nil)
	if err != nil {
		// panic("error creating request: " + err.Error())
		return nil, errors.New(ErrCannotCreateRequest.Error() + ": " + err.Error())
	}

	// @TODO
	if _, ok := f.config.Body.(string); ok {
		req.Header.Set("Content-Type", "text/plain")
	}

	for k, v := range f.config.Headers {
		req.Header.Set(k, v)
	}

	query := req.URL.Query()
	// apply origin query
	for k, v := range urlQueryOrigin {
		query.Add(k, v[0])
	}
	// apply custom query
	for k, v := range f.config.Query {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	// if GET, ignore Body
	if f.config.Body != nil && f.config.Method == GET {
		// // panic("Cannot set body for GET request")
		// return nil, ErrCannotSendBodyWithGet
		f.config.Body = nil
	}

	if f.config.Body != nil {
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}

		if strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			body, err := json.Marshal(f.config.Body)
			if err != nil {
				// panic("error marshalling body: " + err.Error())
				return nil, errors.New(ErrInvalidJSONBody.Error() + ": " + err.Error())
			}

			// req.Header.Set("Content-Type", "application/json")
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		} else if strings.Contains(req.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			body := url.Values{}
			if kv, ok := f.config.Body.(map[string]string); ok {
				for k, v := range kv {
					body.Add(k, v)
				}
			} else {
				return nil, errors.New(ErrInvalidUrlFormEncodedBody.Error() + ": must be map[string]string")
			}

			// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// req.Body = ioutil.NopCloser(bytes.NewReader(body))
			req.Body = ioutil.NopCloser(strings.NewReader(body.Encode()))
		} else if strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") {
			if values, ok := f.config.Body.(map[string]interface{}); ok {
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

					if file, ok := v.(*os.File); ok {
						if fw, err = w.CreateFormFile(k, file.Name()); err != nil {
							return nil, err
						}

						if _, err = io.Copy(fw, file); err != nil {
							return nil, err
						}
						continue
					}
				}
				w.Close()
				req.Header.Set("Content-Type", w.FormDataContentType())
				req.Body = ioutil.NopCloser(&b)
			} else if values, ok := f.config.Body.(map[string]string); ok {
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
				req.Header.Set("Content-Type", w.FormDataContentType())
				req.Body = ioutil.NopCloser(&b)
			} else {
				return nil, errors.New(ErrInvalidBodyMultipart.Error() + ": must be map[string]interface{} or map[string]string")
			}
		} else {
			if _, ok := f.config.Body.(string); !ok {
				return nil, ErrorInvalidBody
			}

			req.Body = ioutil.NopCloser(bytes.NewReader([]byte(f.config.Body.(string))))
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		// panic("error sending request: " + err.Error())
		return nil, errors.New(ErrSendingRequest.Error() + ": " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic("error reading response: " + err.Error())
		return nil, errors.New(ErrReadingResponse.Error() + ": " + err.Error())
	}

	// fmt.Println("response: ", string(body))

	return &Response{
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Body:    body,
	}, nil
}

func (f *Fetch) Send() (*Response, error) {
	return f.Execute()
}

func (f *Fetch) Clone() *Fetch {
	return New(f.config)
}

func (f *Fetch) Head(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(HEAD).
		SetUrl(url)
}

func (f *Fetch) Get(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(GET).
		SetUrl(url)
}

func (f *Fetch) Post(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(POST).
		SetUrl(url)
}

func (f *Fetch) Put(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(PUT).
		SetUrl(url)
}

func (f *Fetch) Patch(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(PATCH).
		SetUrl(url)
}

func (f *Fetch) Delete(url string, config ...*Config) *Fetch {
	return f.Clone().
		SetConfig(config...).
		SetMethod(DELETE).
		SetUrl(url)
}

// func (f *Fetch) JSON() *Response {
// 	f.SetHeader("accept", "application/json")
// 	return f.Execute()
// }
