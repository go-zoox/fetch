package fetch

func Delete(url string, config *Config) *Response {
	return New().Delete(url, config).Execute()
}
