package fetch

func DefaultConfig() *Config {
	config := &Config{
		Headers: make(Headers),
		Query:   make(Query),
		Params:  make(Params),
	}

	config.BaseURL = BaseURL
	config.Timeout = Timeout
	config.Headers["user-agent"] = DefaultUserAgent()

	return config
}

func DefaultUserAgent() string {
	return UserAgent
}
