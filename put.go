package fetch

func Put(url string, config *Config) (*Response, error) {
	return New().Put(url, config).Send()
}
