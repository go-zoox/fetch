package fetch

func Patch(url string, config *Config) (*Response, error) {
	return New().Patch(url, config).Execute()
}
