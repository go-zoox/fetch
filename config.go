package fetch

import (
	"strings"
	"time"
)

type Config struct {
	Url     string
	Method  string
	Headers Headers
	Query   Query
	Params  Params
	Body    Body
	//
	BaseURL string
	Timeout time.Duration
}

func (cfg *Config) Merge(config *Config) {
	// if config == nil {
	// 	return
	// }

	if config.Url != "" {
		cfg.Url = config.Url
	}

	if config.Method != "" {
		cfg.Method = config.Method
	}

	if config.Headers != nil {
		for header := range config.Headers {
			if _, ok := cfg.Headers[header]; !ok {
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

type Body interface{}

type Headers map[string]string

func (h Headers) Get(key string) string {
	for k, v := range h {
		if strings.ToLower(k) == strings.ToLower(key) {
			return strings.ToLower(v)
		}
	}

	return ""
}

type Query map[string]string

func (h Query) Get(key string) string {
	return h[key]
}

type Params map[string]string

func (h Params) Get(key string) string {
	return h[key]
}
