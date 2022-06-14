package fetch

// DefaultConfig returns the default config
func DefaultConfig() *Config {
	config := &Config{
		Headers: make(ConfigHeaders),
		Query:   make(ConfigQuery),
		Params:  make(ConfigParams),
		BaseURL: BaseURL,
		Timeout: Timeout,
	}

	config.Headers["user-agent"] = DefaultUserAgent()

	return config
}

// DefaultUserAgent returns the default user agent
func DefaultUserAgent() string {
	return UserAgent
}
