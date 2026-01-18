# Authentication Examples

## Basic Authentication

```go
response, err := fetch.Get("https://httpbin.zcorky.com/basic-auth/user/pass", &fetch.Config{
	BasicAuth: &fetch.BasicAuth{
		Username: "user",
		Password: "pass",
	},
})
```

### Using Fetch Instance

```go
f := fetch.New()
f.SetBasicAuth("user", "pass")

response, err := f.Get("https://httpbin.zcorky.com/basic-auth/user/pass").Execute()
```

## Bearer Token Authentication

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer your-token-here",
	},
})
```

### Using SetBearerToken

```go
f := fetch.New()
f.SetBearerToken("your-token-here")

response, err := f.Get("https://api.example.com/protected").Execute()
```

## Custom Authorization Header

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Custom token-here",
	},
})
```

## Complete API Client Example

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
	
	// Get all users
	response, err := client.GetUsers()
	if err != nil {
		panic(err)
	}
	
	if response.Ok() {
		fmt.Println(response.JSON())
	}
}
```
