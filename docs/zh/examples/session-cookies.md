# 会话与 Cookie 示例

## 使用 Session

Session 在多个请求之间维护 cookie：

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	// 创建一个 session
	session := fetch.Session()
	session.SetBaseURL("https://httpbin.zcorky.com")

	// 第一个请求 - 登录（设置 cookie）
	loginResp, err := session.Post("/cookies/set", &fetch.Config{
		Body: map[string]interface{}{
			"session": "abc123",
		},
	}).Execute()
	if err != nil {
		panic(err)
	}

	// 后续请求将使用登录时的 cookie
	response, err := session.Get("/cookies").Execute()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## 手动设置 Cookie

```go
f := fetch.New()
f.SetCookie("session", "abc123")
f.SetCookie("user_id", "42")

response, err := f.Get("https://httpbin.zcorky.com/cookies").Execute()
```

## 从响应中获取 Cookie

```go
response, err := fetch.Get("https://httpbin.zcorky.com/cookies/set?name=value")
if err != nil {
	panic(err)
}

// 获取 Set-Cookie 请求头
setCookie := response.SetCookie()
fmt.Println("Set-Cookie:", setCookie)
```

## Session 多个请求

```go
func main() {
	session := fetch.Session()
	session.SetBaseURL("https://api.example.com")

	// 登录
	loginResp, err := session.Post("/login", &fetch.Config{
		Body: map[string]interface{}{
			"username": "user",
			"password": "pass",
		},
	}).Execute()
	
	if err != nil || !loginResp.Ok() {
		panic("登录失败")
	}

	// 受保护的请求 1 - cookie 会自动包含
	usersResp, err := session.Get("/users").Execute()
	
	// 受保护的请求 2 - 同一个 session，相同的 cookie
	profileResp, err := session.Get("/profile").Execute()
	
	fmt.Println(usersResp.JSON())
	fmt.Println(profileResp.JSON())
}
```
