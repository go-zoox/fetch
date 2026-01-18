# Examples

This section provides practical examples for common use cases.

## Overview

- [Basic Usage](./basic) - Simple GET and POST requests
- [HTTP Methods](./http-methods) - All HTTP methods (GET, POST, PUT, PATCH, DELETE, HEAD)
- [Authentication](./auth) - Basic auth, Bearer tokens, custom headers
- [File Operations](./file-operations) - Upload and download files
- [Timeout & Retry](./timeout-retry) - Setting timeouts and retrying failed requests
- [Proxy](./proxy) - Using HTTP, HTTPS, and SOCKS5 proxies
- [Stream](./stream) - Streaming response data
- [Session & Cookies](./session-cookies) - Session management and cookie handling
- [Context Cancellation](./context-cancel) - Canceling requests with context
- [Error Handling](./error-handling) - Handling errors and response status codes

## Quick Example

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/get")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```
