package fetch

// Put is a wrapper for the Put method of the Client
func Put(url string, config *Config) (*Response, error) {
	return New().Put(url, config).Execute()
}
