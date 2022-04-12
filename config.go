package fetch

import (
	"strings"
	"time"
)

type Config struct {
	Url     string
	Method  string
	Headers ConfigHeaders
	Query   ConfigQuery
	Params  ConfigParams
	Body    ConfigBody
	//
	BaseURL string
	Timeout time.Duration
	//
	RetryTimes    int
	RetryInterval time.Duration
}

func (cfg *Config) Merge(config *Config) {
	if config == nil {
		return
	}

	if config.Url != "" {
		cfg.Url = config.Url
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

type ConfigBody interface{}

type ConfigHeaders map[string]string

func (h ConfigHeaders) Get(key string) string {
	for k, v := range h {
		if strings.ToLower(k) == strings.ToLower(key) {
			return strings.ToLower(v)
		}
	}

	return ""
}

type ConfigQuery map[string]string

func (h ConfigQuery) Get(key string) string {
	return h[key]
}

type ConfigParams map[string]string

func (h ConfigParams) Get(key string) string {
	return h[key]
}
