package fetch

func Post(url string, config *Config) (*Response, error) {
	return New().Post(url, config).Execute()
}
