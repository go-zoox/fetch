package fetch

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/go-zoox/headers"
)

// Fetch is the Fetch Client
type Fetch struct {
	config *Config
	Errors []error
}

// New creates a fetch client
func New(cfg ...*Config) *Fetch {
	config := DefaultConfig()
	if len(cfg) > 1 {
		panic("Too many arguments")
	}

	if len(cfg) == 1 {
		config.Merge(cfg[0])
	}

	if config.Context == nil {
		config.Context = context.Background()
	}

	return &Fetch{
		config: config,
	}
}

// Create creates a new fetch with base url
// Specially useful for Client SDK
func Create(baseURL string) *Fetch {
	return New().SetBaseURL(baseURL)
}

// SetContext sets the context
func (f *Fetch) SetContext(ctx context.Context) *Fetch {
	f.config.Context = ctx
	return f
}

// SetConfig sets the config of fetch
func (f *Fetch) SetConfig(configs ...*Config) *Fetch {
	for _, config := range configs {
		f.config.Merge(config)
	}

	return f
}

// SetURL sets the url of fetch
func (f *Fetch) SetURL(url string) *Fetch {
	f.config.URL = url
	return f
}

// SetDownloadFilePath sets the download file path
func (f *Fetch) SetDownloadFilePath(filepath string) *Fetch {
	f.config.DownloadFilePath = filepath
	return f
}

// SetProgressCallback sets the progress callback
func (f *Fetch) SetProgressCallback(callback func(percent int64, current, total int64)) *Fetch {
	f.config.OnProgress = &callback
	return f
}

// SetMethod sets the method
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

// SetHeader sets the header key and value
func (f *Fetch) SetHeader(key, value string) *Fetch {
	f.config.Headers[key] = value
	return f
}

// SetQuery sets the query key and value
func (f *Fetch) SetQuery(key, value string) *Fetch {
	f.config.Query[key] = value
	return f
}

// SetParam sets the param key and value
func (f *Fetch) SetParam(key, value string) *Fetch {
	f.config.Params[key] = value
	return f
}

// SetBody sets the body
func (f *Fetch) SetBody(body Body) *Fetch {
	f.config.Body = body
	return f
}

// SetBaseURL sets the base url
func (f *Fetch) SetBaseURL(url string) *Fetch {
	f.config.BaseURL = url
	return f
}

// SetTimeout sets the timeout
func (f *Fetch) SetTimeout(timeout time.Duration) *Fetch {
	f.config.Timeout = timeout
	return f
}

// SetUserAgent sets the user agent
func (f *Fetch) SetUserAgent(userAgent string) *Fetch {
	return f.SetHeader(headers.UserAgent, userAgent)
}

// SetBasicAuth sets the basic auth username and password
func (f *Fetch) SetBasicAuth(username, password string) *Fetch {
	return f.SetAuthorization("Basic " + base64.RawStdEncoding.EncodeToString([]byte(username+":"+password)))
}

// SetBearToken sets the bear token
func (f *Fetch) SetBearToken(token string) *Fetch {
	return f.SetAuthorization("Bearer " + token)
}

// SetAuthorization sets the authorization token
func (f *Fetch) SetAuthorization(token string) *Fetch {
	return f.SetHeader(headers.Authorization, token)
}

// SetCookie sets the cookie
func (f *Fetch) SetCookie(key, value string) *Fetch {
	origin := f.config.Headers.Get(headers.Cookie)

	cookie := newCookie(origin)
	cookie.Set(key, value)

	return f.SetHeader(headers.Cookie, cookie.String())
}

// SetProxy sets the proxy
//
//		support http, https, socks5
//	 example:
//			http://127.0.0.1:17890
//		  https://127.0.0.1:17890
//		  socks5://127.0.0.1:17890
func (f *Fetch) SetProxy(proxy string) *Fetch {
	// validdate proxy
	_, err := url.Parse(proxy)
	if err != nil {
		panic(fmt.Sprintf("invalid proxy %s", proxy))
	}

	f.config.Proxy = proxy

	return f
}

// SetAccept sets the accept header
func (f *Fetch) SetAccept(accept string) *Fetch {
	return f.SetHeader(headers.Accept, accept)
}

// SetContentType ...
func (f *Fetch) SetContentType(contentType string) *Fetch {
	return f.SetHeader(headers.ContentType, contentType)
}

// SetReferrer sets the referrer
func (f *Fetch) SetReferrer(referrer string) *Fetch {
	return f.SetHeader(headers.Referrer, referrer)
}

// SetCacheControl sets the cache control
func (f *Fetch) SetCacheControl(cacheControl string) *Fetch {
	return f.SetHeader(headers.CacheControl, cacheControl)
}

// SetAcceptEncoding sets the accept encoding
func (f *Fetch) SetAcceptEncoding(acceptEncoding string) *Fetch {
	return f.SetHeader(headers.AcceptEncoding, acceptEncoding)
}

// SetAcceptLanguage sets the accept language
func (f *Fetch) SetAcceptLanguage(acceptLanguage string) *Fetch {
	return f.SetHeader(headers.AcceptLanguage, acceptLanguage)
}

// Config returns the config of fetch
func (f *Fetch) Config() (*Config, error) {
	cfg := f.config.Clone()

	// if f.isConfigBuilt {
	// 	return f.config, nil
	// }
	// f.isConfigBuilt = true

	newURL := f.config.URL
	if f.config.Params != nil {
		for k, v := range f.config.Params {
			vEscaped := url.QueryEscape(v)
			// support /:id/:name
			newURL = strings.Replace(newURL, ":"+k, vEscaped, -1)
			// support /{id}/{name}
			newURL = strings.Replace(newURL, "{"+k+"}", vEscaped, -1)
		}
	}

	// @BASEURL
	if f.config.BaseURL != "" {
		uNewURL, err := url.Parse(newURL)
		if err != nil {
			return cfg, errors.New("invalid NewURL")
		}

		if uNewURL.Host == "" {
			parsedBaseURL, err := url.Parse(f.config.BaseURL)
			if err != nil {
				return cfg, errors.New("invalid base URL")
			}

			parsedBaseURL.Path = path.Join(parsedBaseURL.Path, newURL)
			newURL = parsedBaseURL.String()
		}
	}

	cfg.URL = newURL

	return cfg, nil
}

// Send sends the request
func (f *Fetch) Send() (*Response, error) {
	return f.Execute()
}

// Clone creates a new fetch
func (f *Fetch) Clone() *Fetch {
	return New(f.config)
}

// Retry retries the request
func (f *Fetch) Retry(before func(f *Fetch)) (*Response, error) {
	nf := f.Clone()

	if before != nil {
		before(nf)
	}

	return nf.Send()
}
