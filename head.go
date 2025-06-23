package fetch

// Head is a wrapper for the Head method of the Client
func Head(url string, config ...interface{}) (*Response, error) {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		return nil, ErrTooManyArguments
	}

	if c.Body != nil {
		panic("Request with HEAD method cannot have body")
	}

	return New().Head(url, c).Execute()
}
