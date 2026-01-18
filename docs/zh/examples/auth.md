# 认证示例

## Basic 认证

```go
response, err := fetch.Get("https://httpbin.zcorky.com/basic-auth/user/pass", &fetch.Config{
	BasicAuth: &fetch.BasicAuth{
		Username: "user",
		Password: "pass",
	},
})
```

### 使用 Fetch 实例

```go
f := fetch.New()
f.SetBasicAuth("user", "pass")

response, err := f.Get("https://httpbin.zcorky.com/basic-auth/user/pass").Execute()
```

## Bearer Token 认证

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer your-token-here",
	},
})
```

### 使用 SetBearerToken

```go
f := fetch.New()
f.SetBearerToken("your-token-here")

response, err := f.Get("https://api.example.com/protected").Execute()
```

## 自定义认证请求头

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Custom token-here",
	},
})
```

## 完整的 API 客户端示例

```go
package main

import (
	"fmt"
	"time"
	"github.com/go-zoox/fetch"
)

type APIClient struct {
	baseURL string
	token   string
	client  *fetch.Fetch
}

func NewAPIClient(baseURL, token string) *APIClient {
	f := fetch.New()
	f.SetBaseURL(baseURL)
	f.SetTimeout(10 * time.Second)
	
	return &APIClient{
		baseURL: baseURL,
		token:   token,
		client:  f,
	}
}

func (c *APIClient) GetUsers() (*fetch.Response, error) {
	return c.client.SetBearerToken(c.token).Get("/users").Execute()
}

func (c *APIClient) GetUser(id string) (*fetch.Response, error) {
	return c.client.SetBearerToken(c.token).Get("/users/" + id).Execute()
}

func (c *APIClient) CreateUser(user map[string]interface{}) (*fetch.Response, error) {
	return c.client.SetBearerToken(c.token).Post("/users", &fetch.Config{
		Body: user,
	}).Execute()
}

func main() {
	client := NewAPIClient("https://api.example.com", "your-api-token")
	
	// 获取所有用户
	response, err := client.GetUsers()
	if err != nil {
		panic(err)
	}
	
	if response.Ok() {
		fmt.Println(response.JSON())
	}
}
```
