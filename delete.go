package fetch

// Delete is a wrapper for the Delete method of the Client
func Delete(url string, config *Config) (*Response, error) {
	return New().Delete(url, config).Execute()
}
