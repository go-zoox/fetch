package fetch

// Patch is a wrapper for the Patch method of the Client
func Patch(url string, config *Config) (*Response, error) {
	return New().Patch(url, config).Execute()
}
