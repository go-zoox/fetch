package fetch

func Put(url string, config *Config) *Response {
	return New().Put(url, config).Execute()
}
