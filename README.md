# Fetch - HTTP Client

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/fetch)](https://pkg.go.dev/github.com/go-zoox/fetch)
[![Build Status](https://github.com/go-zoox/fetch/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/fetch/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/fetch)](https://goreportcard.com/report/github.com/go-zoox/fetch)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/fetch/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/fetch?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/fetch.svg)](https://github.com/go-zoox/fetch/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/fetch.svg?label=Release)](https://github.com/go-zoox/fetch/releases)

Load application environment variables from a `.env` file into the current process.

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/fetch
```

## Getting Started

```go
response, _ := fetch.Get("https://httpbin.zcorky.com/get")
url := response.Get("url")
method := response.Get("method")

fmt.Println(url, method)
```

## License
GoZoox is released under the [MIT License](./LICENSE).