# Installation

## Requirements

- Go 1.18 or higher

## Install

Install the package using Go modules:

```bash
go get github.com/go-zoox/fetch
```

## Import

Import the package in your Go code:

```go
import "github.com/go-zoox/fetch"
```

## Verify Installation

You can verify the installation by creating a simple test file:

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
	fmt.Println("Installation successful!", response.Status)
}
```

Run it:

```bash
go run main.go
```

If you see "Installation successful!" and a status code, you're all set!
