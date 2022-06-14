package fetch

// Post is a wrapper for the Post method of the Client
func Post(url string, config *Config) (*Response, error) {
	return New().Post(url, config).Execute()
}
