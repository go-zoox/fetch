package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Fetch struct {
	config *Config
	Errors []error
}

func New() *Fetch {
	config := DefaultConfig()

	return &Fetch{config: config}
}

func (f *Fetch) SetConfig(config *Config) *Fetch {
	for header := range config.Headers {
		if _, ok := f.config.Headers[header]; !ok {
			f.SetHeader(header, config.Headers[header])
		}
	}

	for query := range config.Query {
		if _, ok := f.config.Query[query]; !ok {
			f.SetQuery(query, config.Query[query])
		}
	}

	for param := range config.Params {
		if _, ok := f.config.Params[param]; !ok {
			f.SetParam(param, config.Params[param])
		}
	}

	if config.Body != nil {
		f.SetBody(config.Body)
	}

	return f
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

func (f *Fetch) SetBody(body Body) *Fetch {
	f.config.Body = body
	return f
}

func (f *Fetch) Execute() (*Response, error) {
	if len(f.Errors) > 0 {
		return nil, f.Errors[0]
	}

	client := &http.Client{}
	req, err := http.NewRequest(f.config.Method, f.config.Url, nil)
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
	if req.Header.Get("content-type") == "" {
		req.Header.Set("content-type", "application/json")
	}

	query := req.URL.Query()
	for k, v := range f.config.Query {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	if f.config.Body != nil {
		if f.config.Method == GET {
			// panic("Cannot set body for GET request")
			return nil, ErrCannotSendBodyWithGet
		}

		if strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			body, err := json.Marshal(f.config.Body)
			if err != nil {
				// panic("error marshalling body: " + err.Error())
				return nil, errors.New(ErrInvalidJSONBody.Error() + ": " + err.Error())
			}

			// req.Header.Set("content-type", "application/json")
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		} else if strings.Contains(req.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			body := url.Values{}
			if kv, ok := f.config.Body.(map[string]string); ok {
				for k, v := range kv {
					body.Add(k, v)
				}
			} else {
				return nil, ErrInvalidUrlFormEncodedBody
			}

			// req.Header.Set("content-type", "application/x-www-form-urlencoded")
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
				req.Header.Set("content-type", w.FormDataContentType())
				req.Body = ioutil.NopCloser(&b)
			} else {
				return nil, ErrInvalidBodyMultipart
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

func (f *Fetch) Get(url string, config *Config) *Fetch {
	f.SetConfig(config)
	f.SetMethod(GET)
	f.SetUrl(url)
	return f
}

func (f *Fetch) Post(url string, config *Config) *Fetch {
	f.SetConfig(config)
	f.SetMethod(POST)
	f.SetUrl(url)
	return f
}

func (f *Fetch) Put(url string, config *Config) *Fetch {
	f.SetConfig(config)
	f.SetMethod(PUT)
	f.SetUrl(url)
	return f
}

func (f *Fetch) Patch(url string, config *Config) *Fetch {
	f.SetConfig(config)
	f.SetMethod(PATCH)
	f.SetUrl(url)
	return f
}

func (f *Fetch) Delete(url string, config *Config) *Fetch {
	f.SetConfig(config)
	f.SetMethod(DELETE)
	f.SetUrl(url)
	return f
}

// func (f *Fetch) JSON() *Response {
// 	f.SetHeader("accept", "application/json")
// 	return f.Execute()
// }
