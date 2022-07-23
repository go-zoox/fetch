package fetch

import (
	"strings"
	"time"
)

// Config is the configuration for the fetch
type Config struct {
	URL     string
	Method  string
	Headers ConfigHeaders
	Query   ConfigQuery
	Params  ConfigParams
	Body    ConfigBody
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
		for header := range config.Headers {
			if _, ok := c.Headers[header]; !ok {
				// fmt.Printf("%s origin(%s) => new(%s)", header, cfg.Headers[header], config.Headers[header])
				c.Headers[header] = config.Headers[header]
			}
		}
	}

	if config.Query != nil {
		for query := range config.Query {
			if _, ok := c.Query[query]; !ok {
				c.Query[query] = config.Query[query]
			}
		}
	}

	if config.Params != nil {
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

// ConfigBody is the body of the request
type ConfigBody interface{}

// ConfigHeaders is the headers of the request
type ConfigHeaders map[string]string

// Get returns the value of the given key
func (h ConfigHeaders) Get(key string) string {
	for k, v := range h {
		if strings.EqualFold(k, key) {
			return strings.ToLower(v)
		}
	}

	return ""
}

// ConfigQuery is the query of the request
type ConfigQuery map[string]string

// Get returns the value of the given key
func (h ConfigQuery) Get(key string) string {
	return h[key]
}

// ConfigParams is the params of the request
type ConfigParams map[string]string

// Get returns the value of the given key
func (h ConfigParams) Get(key string) string {
	return h[key]
}
