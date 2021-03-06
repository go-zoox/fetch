package fetch

// Download is a wrapper for the Download method of the Client
func Download(url string, filepath string, config ...interface{}) (*Response, error) {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		return nil, ErrTooManyArguments
	}

	return New().Download(url, filepath, c).Execute()
}
