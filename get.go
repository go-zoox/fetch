package fetch

func Get(url string, config ...interface{}) *Response {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		panic("too many arguments")
	}

	return New().Get(url, c).Execute()
}
