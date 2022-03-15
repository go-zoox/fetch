package fetch

import "fmt"

func DefaultConfig() *Config {
	config := &Config{
		Headers: make(ConfigHeaders),
		Query:   make(ConfigQuery),
		Params:  make(ConfigParams),
	}

	config.Headers["user-agent"] = DefaultUserAgent()

	return config
}

func DefaultUserAgent() string {
	return fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version)
}
