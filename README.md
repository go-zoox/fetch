# Fetch - HTTP Client

`HTTP Client` for Go, inspired by the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API), and [Axios](https://github.com/axios/axios) + [Got](https://github.com/sindresorhus/got) (Sindre Sorhus).

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/fetch)](https://pkg.go.dev/github.com/go-zoox/fetch)
[![Build Status](https://github.com/go-zoox/fetch/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/fetch/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/fetch)](https://goreportcard.com/report/github.com/go-zoox/fetch)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/fetch/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/fetch?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/fetch.svg)](https://github.com/go-zoox/fetch/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/fetch.svg?label=Release)](https://github.com/go-zoox/fetch/releases)

## Features
### Main API
- [x] Make HTTP requests
- [x] Easy JSON Response
- [ ] GZip support
  - [x] Decode GZip response
  - [ ] Encode GZip request (Upload File with GZip)
- [ ] HTTP/2 support
- [ ] Simple Auth Methods
  - [x] Basic Auth
  - [x] Bearer Auth
- [ ] Support cancel

### Timeouts and retries
- [x] Support timeout
- [x] Support retry on failure

### Progress
- [ ] Support progress and progress events

### File upload and download
- [x] Download files easily
- [x] Upload files easily

### Cache, Proxy and UNIX sockets
- [ ] [RFC compliant caching](https://github.com/sindresorhus/got/blob/main/documentation/cache.md)
- [x] Proxy support
  - [x] Environment variables (HTTP_PROXY/HTTPS_PROXY/SOCKS_PROXY)
  - [x] Custom proxy
- [ ] UNIX Domain Sockets

### WebDAV
- [ ] WebDAV protocol support

### Advanced creation
- [ ] Plugin system
- [ ] Middleware system

## Methods


## Installation

To install the package, run:

```bash
go get github.com/go-zoox/fetch
```

## Methods
- [x] GET
- [x] POST
- [x] PUT
- [x] PATCH
- [x] DELETE
- [x] HEAD
- [ ] OPTIONS
- [ ] TRACE
- [ ] CONNECT

## Getting Started

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
  response, _ := fetch.Get("https://httpbin.zcorky.com/get")
  url := response.Get("url")
  method := response.Get("method")

  fmt.Println(url, method)
}
```

## Examples

### Get

```go
package main

import (
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

### Post

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
  response, _ := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
		Body: map[string]interface{}{
			"foo":     "bar",
			"foo2":    "bar2",
			"number":  1,
			"boolean": true,
			"array": []string{
				"foo3",
				"bar3",
			},
			"nest": map[string]string{
				"foo4": "bar4",
			},
		},
	})
	if err != nil {
		panic(err)
	}

  fmt.Println(response.JSON())
}
```

### Put

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
		Body: map[string]interface{}{
			"foo":     "bar",
			"foo2":    "bar2",
			"number":  1,
			"boolean": true,
			"array": []string{
				"foo3",
				"bar3",
			},
			"nest": map[string]string{
				"foo4": "bar4",
			},
		},
	})
  if err != nil {
		panic(err)
	}


  fmt.Println(response.JSON())
}
```
### Delete

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Delete("https://httpbin.zcorky.com/Delete", &fetch.Config{
		Body: map[string]interface{}{
			"foo":     "bar",
			"foo2":    "bar2",
			"number":  1,
			"boolean": true,
			"array": []string{
				"foo3",
				"bar3",
			},
			"nest": map[string]string{
				"foo4": "bar4",
			},
		},
	})
	if err != nil {
		panic(err)
	}

  fmt.Println(response.JSON())
}
```

### Timeout

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
  response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
    Timeout: 5 * time.Second,
  })
  if err != nil {
		panic(err)
	}

  fmt.Println(response.JSON())
}
```

### Proxy

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
  response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
    Proxy: "http://127.0.0.1:17890",
  })
  if err != nil {
		panic(err)
	}

  fmt.Println(response.JSON())
}
```

### Basic Auth

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
  response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
    BasicAuth: &fetch.BasicAuth{
      Username: "foo",
      Password: "bar",
    },
  })
  if err != nil {
		panic(err)
	}

  fmt.Println(response.JSON())
}
```

### Download

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Download("https://httpbin.zcorky.com/image", "/tmp/image.webp")
  if err != nil {
		panic(err)
	}
}
```

### Upload

```go
package main

import (
  "github.com/go-zoox/fetch"
)

func main() {
		file, _ := os.Open("go.mod")

	response, err := Upload("https://httpbin.zcorky.com/upload", file)
  if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## Depencencies

- [gjson](github.com/tidwall/gjson) - Get JSON Whenever You Need, you don't
  define type first。

## Inspired By

- [sindresorhus/got](https://github.com/sindresorhus/got) - 🌐 Human-friendly and powerful HTTP request library for Node.js
- [axios/axios](https://github.com/axios/axios) - Promise based HTTP client for the browser and node.js
- [mozillazg/request](https://github.com/mozillazg/request) - A
  developer-friendly HTTP request library for Gopher
- [monaco-io/request](https://github.com/monaco-io/request) - go request, go http client
## License

GoZoox is released under the [MIT License](./LICENSE).
