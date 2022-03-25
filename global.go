package fetch

import (
	"fmt"
	"time"
)

var BaseURL = ""
var Timeout = 60 * time.Second
var UserAgent = fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version)

// @TODO
// var Headers = make(ConfigHeaders)

func SetBaseURL(url string) {
	BaseURL = url
}

func SetTimeout(timeout time.Duration) {
	Timeout = timeout
}

func SetUserAgent(userAgent string) {
	UserAgent = userAgent
}

// func SetHeader(key, value string) {
// 	Headers[key] = value
// }
