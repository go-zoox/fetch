package fetch

// Stream is a wrapper for the Stream method of the Client
func Stream(url string, config ...interface{}) (*Response, error) {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		return nil, ErrTooManyArguments
	}

	return New().Stream(url, c).Execute()
}
