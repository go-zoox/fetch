package fetch

import "io"

// Upload is a wrapper for the Upload method of the Client
func Upload(url string, file io.Reader, config ...interface{}) (*Response, error) {
	c := &Config{}
	if len(config) == 1 {
		c = config[0].(*Config)
	} else if len(config) > 1 {
		return nil, ErrTooManyArguments
	}

	return New().Upload(url, file, c).Execute()
}
