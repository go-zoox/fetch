# Proxy

Fetch supports HTTP, HTTPS, and SOCKS5 proxies.

## Basic Proxy Usage

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "http://127.0.0.1:8080",
})
```

## Proxy Types

### HTTP Proxy

```go
&fetch.Config{
	Proxy: "http://127.0.0.1:8080",
}
```

### HTTPS Proxy

```go
&fetch.Config{
	Proxy: "https://127.0.0.1:8080",
}
```

### SOCKS5 Proxy

```go
&fetch.Config{
	Proxy: "socks5://127.0.0.1:1080",
}
```

### Proxy with Authentication

```go
&fetch.Config{
	Proxy: "http://user:password@127.0.0.1:8080",
}
```

## Environment Variables

Fetch automatically uses proxy settings from environment variables:

- `HTTP_PROXY`
- `HTTPS_PROXY`
- `SOCKS_PROXY`

You don't need to configure them explicitly if environment variables are set.

```bash
export HTTP_PROXY=http://127.0.0.1:8080
export HTTPS_PROXY=http://127.0.0.1:8080
```

## Using Fetch Instance

```go
f := fetch.New()
f.SetProxy("http://127.0.0.1:8080")

response, err := f.Get("https://httpbin.zcorky.com/ip").Execute()
```

## Example

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	// Use proxy
	response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
		Proxy: "http://127.0.0.1:17890",
	})
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.JSON())
}
```
