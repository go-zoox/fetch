package fetch

import (
	"context"
	"strings"
	"time"
)

// Config is the configuration for the fetch
type Config struct {
	URL     string
	Method  string
	Headers Headers
	Query   Query
	Params  Params
	Body    Body
	//
	BaseURL string
	Timeout time.Duration
	//
	DownloadFilePath string
	//
	Proxy string
	//
	IsStream bool
	//
	IsSession bool
	//
	HTTP2 bool
	//
	TLSCertificate     []byte
	TLSCertificateFile string
	//
	Context context.Context
	//
	OnProgress func(percent int64, current, total int64)
}

// Merge merges the config with the given config
func (c *Config) Merge(config *Config) {
	if config == nil {
		return
	}

	if config.URL != "" {
		c.URL = config.URL
	}

	if config.Method != "" {
		c.Method = config.Method
	}

	if config.Headers != nil {
		if c.Headers == nil {
			c.Headers = make(Headers)
		}

		for header := range config.Headers {
			if _, ok := c.Headers[header]; !ok {
				// fmt.Printf("%s origin(%s) => new(%s)", header, cfg.Headers[header], config.Headers[header])
				c.Headers[header] = config.Headers[header]
			}
		}
	}

	if config.Query != nil {
		if c.Query == nil {
			c.Query = make(Query)
		}

		for query := range config.Query {
			if _, ok := c.Query[query]; !ok {
				c.Query[query] = config.Query[query]
			}
		}
	}

	if config.Params != nil {
		if c.Params == nil {
			c.Params = make(Params)
		}

		for param := range config.Params {
			if _, ok := c.Params[param]; !ok {
				c.Params[param] = config.Params[param]
			}
		}
	}

	if config.Body != nil {
		c.Body = config.Body
	}

	if config.BaseURL != "" {
		c.BaseURL = config.BaseURL
	}

	if config.Timeout != 0 {
		c.Timeout = config.Timeout
	}

	if config.DownloadFilePath != "" {
		c.DownloadFilePath = config.DownloadFilePath
	}

	if config.Proxy != "" {
		c.Proxy = config.Proxy
	}

	if config.TLSCertificate != nil {
		c.TLSCertificate = config.TLSCertificate
	}

	if config.TLSCertificateFile != "" {
		c.TLSCertificateFile = config.TLSCertificateFile
	}

	if config.IsStream {
		c.IsStream = config.IsStream
	}
}

// Clone returns a clone of the config
func (c *Config) Clone() *Config {
	nc := DefaultConfig()
	nc.Merge(c)
	return c
}

// Body is the body of the request
type Body interface{}

// Headers is the headers of the request
type Headers map[string]string

// Get returns the value of the given key
func (h Headers) Get(key string) string {
	for k, v := range h {
		if strings.EqualFold(k, key) {
			return v
		}
	}

	return ""
}

// Set sets the value of the given key
func (h Headers) Set(key, value string) {
	h[strings.ToLower(key)] = value
}

// Query is the query of the request
type Query map[string]string

// Get returns the value of the given key
func (q Query) Get(key string) string {
	return q[key]
}

// Set sets the value of the given key
func (q Query) Set(key, value string) {
	q[key] = value
}

// Params is the params of the request
type Params map[string]string

// Get returns the value of the given key
func (p Params) Get(key string) string {
	return p[key]
}

// Set sets the value of the given key
func (p Params) Set(key, value string) {
	p[key] = value
}
