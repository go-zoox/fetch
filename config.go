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
}

// Merge merges the config with the given config
func (cfg *Config) Merge(config *Config) {
	if config == nil {
		return
	}

	if config.URL != "" {
		cfg.URL = config.URL
	}

	if config.Method != "" {
		cfg.Method = config.Method
	}

	if config.Headers != nil {
		for header := range config.Headers {
			if _, ok := cfg.Headers[header]; !ok {
				// fmt.Printf("%s origin(%s) => new(%s)", header, cfg.Headers[header], config.Headers[header])
				cfg.Headers[header] = config.Headers[header]
			}
		}
	}

	if config.Query != nil {
		for query := range config.Query {
			if _, ok := cfg.Query[query]; !ok {
				cfg.Query[query] = config.Query[query]
			}
		}
	}

	if config.Params != nil {
		for param := range config.Params {
			if _, ok := cfg.Params[param]; !ok {
				cfg.Params[param] = config.Params[param]
			}
		}
	}

	if config.Body != nil {
		cfg.Body = config.Body
	}

	if config.BaseURL != "" {
		cfg.BaseURL = config.BaseURL
	}

	if config.Timeout != 0 {
		cfg.Timeout = config.Timeout
	}
}

// ConfigBody is the body of the request
type ConfigBody interface{}

// ConfigHeaders is the headers of the request
type ConfigHeaders map[string]string

// Get returns the value of the given key
func (h ConfigHeaders) Get(key string) string {
	for k, v := range h {
		if strings.ToLower(k) == strings.ToLower(key) {
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
