package fetch

func Post(url string, config *Config) *Response {
	return New().Post(url, config).Execute()
}
