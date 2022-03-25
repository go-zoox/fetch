package fetch

import (
	"fmt"
	"time"
)

var BaseURL = ""
var Timeout = 60 * time.Second
var UserAgent = fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version)
