package fetch

func Delete(url string, config *Config) (*Response, error) {
	return New().Delete(url, config).Execute()
}
