# HTTP 方法示例

所有支持的 HTTP 方法的示例：GET、POST、PUT、PATCH、DELETE 和 HEAD。

## GET 请求

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

## 带查询参数的 GET

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Query: map[string]string{
		"page": "1",
		"limit": "10",
	},
})
```

## POST 请求

```go
response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PUT 请求

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
		"id": 1,
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PATCH 请求

```go
response, err := fetch.Patch("https://httpbin.zcorky.com/patch", &fetch.Config{
	Body: map[string]interface{}{
		"email": "newemail@example.com",
	},
})
```

## DELETE 请求

```go
response, err := fetch.Delete("https://httpbin.zcorky.com/delete", &fetch.Config{
	Body: map[string]interface{}{
		"id": 1,
	},
})
```

## HEAD 请求

```go
response, err := fetch.Head("https://httpbin.zcorky.com/get")
if err != nil {
	panic(err)
}

// 检查请求头而不下载响应体
contentType := response.ContentType()
contentLength := response.ContentLength()
```

## 使用 Fetch 实例

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

## 方法链式调用

```go
response, err := fetch.New().
	SetBaseURL("https://api.example.com").
	SetTimeout(10*time.Second).
	SetHeader("Authorization", "Bearer token").
	Get("/users").
	Execute()
```
