package fetch

func Patch(url string, config *Config) *Response {
	return New().Patch(url, config).Execute()
}
