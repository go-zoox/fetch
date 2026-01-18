# 基础用法示例

## 简单 GET 请求

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
		"foo": "bar",
		"number": 42,
		"boolean": true,
		"array": []string{"item1", "item2"},
		"nest": map[string]string{
			"key": "value",
		},
	},
})
```

## PUT 请求

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
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
response, err := fetch.Delete("https://httpbin.zcorky.com/delete")
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

## 处理响应

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	// 访问 JSON 值
	name := response.Get("name")
	email := response.Get("email")
	
	// 或反序列化为结构体
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	var user User
	err := response.UnmarshalJSON(&user)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("用户: %s (%s)\n", user.Name, user.Email)
} else {
	fmt.Printf("错误: %v\n", response.Error())
}
```
