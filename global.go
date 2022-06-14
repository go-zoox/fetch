package fetch

import (
	"fmt"
	"time"
)

// BaseURL is the default base url
var BaseURL = ""

// Timeout is the default timeout
var Timeout = 60 * time.Second

// UserAgent is the default user agent
var UserAgent = fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version)

// @TODO
// var Headers = make(ConfigHeaders)

// SetBaseURL sets the base url
func SetBaseURL(url string) {
	BaseURL = url
}

// SetTimeout sets the timeout
func SetTimeout(timeout time.Duration) {
	Timeout = timeout
}

// SetUserAgent sets the user agent
func SetUserAgent(userAgent string) {
	UserAgent = userAgent
}

// func SetHeader(key, value string) {
// 	Headers[key] = value
// }
