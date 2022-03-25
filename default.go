package fetch

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

func DefaultUserAgent() string {
	return UserAgent
}
