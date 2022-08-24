package fetch

// DefaultConfig returns the default config
func DefaultConfig() *Config {
	config := &Config{
		Headers: make(Headers),
		Query:   make(Query),
		Params:  make(Params),
		BaseURL: BaseURL,
		Timeout: Timeout,
	}

	config.Headers[HeaderUserAgent] = DefaultUserAgent()

	return config
}

// DefaultUserAgent returns the default user agent
func DefaultUserAgent() string {
	return UserAgent
}
