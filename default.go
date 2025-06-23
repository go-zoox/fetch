package fetch

import "github.com/go-zoox/headers"

// DefaultConfig returns the default config
func DefaultConfig() *Config {
	config := &Config{
		Headers:         make(Headers),
		Query:           make(Query),
		Params:          make(Params),
		BaseURL:         BaseURL,
		Timeout:         Timeout,
		CompressRequest: false,
	}

	config.Headers[headers.UserAgent] = DefaultUserAgent()

	return config
}

// DefaultUserAgent returns the default user agent
func DefaultUserAgent() string {
	return UserAgent
}
