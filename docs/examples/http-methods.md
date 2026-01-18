# HTTP Methods Examples

Examples for all supported HTTP methods: GET, POST, PUT, PATCH, DELETE, and HEAD.

## GET Request

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

## GET with Query Parameters

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Query: map[string]string{
		"page": "1",
		"limit": "10",
	},
})
```

## POST Request

```go
response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PUT Request

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
		"id": 1,
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PATCH Request

```go
response, err := fetch.Patch("https://httpbin.zcorky.com/patch", &fetch.Config{
	Body: map[string]interface{}{
		"email": "newemail@example.com",
	},
})
```

## DELETE Request

```go
response, err := fetch.Delete("https://httpbin.zcorky.com/delete", &fetch.Config{
	Body: map[string]interface{}{
		"id": 1,
	},
})
```

## HEAD Request

```go
response, err := fetch.Head("https://httpbin.zcorky.com/get")
if err != nil {
	panic(err)
}

// Check headers without downloading body
contentType := response.ContentType()
contentLength := response.ContentLength()
```

## Using Fetch Instance

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")

// GET
response, err := f.Get("/users").Execute()

// POST
response, err := f.Post("/users", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
	},
}).Execute()

// PUT
response, err := f.Put("/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"name": "Jane",
	},
}).Execute()

// PATCH
response, err := f.Patch("/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"email": "new@example.com",
	},
}).Execute()

// DELETE
response, err := f.Delete("/users/1").Execute()

// HEAD
response, err := f.Head("/users/1").Execute()
```

## Method Chaining

```go
response, err := fetch.New().
	SetBaseURL("https://api.example.com").
	SetTimeout(10*time.Second).
	SetHeader("Authorization", "Bearer token").
	Get("/users").
	Execute()
```
