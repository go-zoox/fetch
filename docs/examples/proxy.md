# Proxy Examples

## HTTP Proxy

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
		Proxy: "http://127.0.0.1:8080",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## HTTPS Proxy

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "https://127.0.0.1:8080",
})
```

## SOCKS5 Proxy

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "socks5://127.0.0.1:1080",
})
```

## Proxy with Authentication

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "http://username:password@127.0.0.1:8080",
})
```

## Using Fetch Instance

```go
f := fetch.New()
f.SetProxy("http://127.0.0.1:8080")

response, err := f.Get("https://httpbin.zcorky.com/ip").Execute()
```

## Environment Variables

Fetch automatically uses proxy settings from environment variables:

```bash
export HTTP_PROXY=http://127.0.0.1:8080
export HTTPS_PROXY=http://127.0.0.1:8080
export SOCKS_PROXY=socks5://127.0.0.1:1080
```

No need to configure explicitly in code:

```go
// Automatically uses HTTP_PROXY/HTTPS_PROXY/SOCKS_PROXY
response, err := fetch.Get("https://httpbin.zcorky.com/ip")
```
