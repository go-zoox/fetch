package fetch

// Get is a wrapper for the Get method of the Client
func Get(url string, config ...interface{}) (*Response, error) {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		return nil, ErrTooManyArguments
	}

	if c.Body != nil {
		panic("Request with GET method cannot have body")
	}

	return New().Get(url, c).Execute()
}
