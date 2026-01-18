# Quick Start

## Installation

Install the package using Go modules:

```bash
go get github.com/go-zoox/fetch
```

## Your First Request

Here's a simple example to get you started:

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

	url := response.Get("url")
	method := response.Get("method")

	fmt.Println(url, method)
}
```

## Basic GET Request

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

## Basic POST Request

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
		Body: map[string]interface{}{
			"foo": "bar",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## Next Steps

- Read the [Guide](/guide/) for detailed documentation
- Check out [Examples](/examples/) for more use cases
- Browse the [API Reference](/api/) for complete API documentation
