# Session & Cookies Examples

## Using Session

Session maintains cookies across multiple requests:

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	// Create a session
	session := fetch.Session()
	session.SetBaseURL("https://httpbin.zcorky.com")

	// First request - login (sets cookie)
	loginResp, err := session.Post("/cookies/set", &fetch.Config{
		Body: map[string]interface{}{
			"session": "abc123",
		},
	}).Execute()
	if err != nil {
		panic(err)
	}

	// Subsequent requests will use cookies from login
	response, err := session.Get("/cookies").Execute()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## Setting Cookies Manually

```go
f := fetch.New()
f.SetCookie("session", "abc123")
f.SetCookie("user_id", "42")

response, err := f.Get("https://httpbin.zcorky.com/cookies").Execute()
```

## Getting Cookies from Response

```go
response, err := fetch.Get("https://httpbin.zcorky.com/cookies/set?name=value")
if err != nil {
	panic(err)
}

// Get Set-Cookie header
setCookie := response.SetCookie()
fmt.Println("Set-Cookie:", setCookie)
```

## Session with Multiple Requests

```go
func main() {
	session := fetch.Session()
	session.SetBaseURL("https://api.example.com")

	// Login
	loginResp, err := session.Post("/login", &fetch.Config{
		Body: map[string]interface{}{
			"username": "user",
			"password": "pass",
		},
	}).Execute()
	
	if err != nil || !loginResp.Ok() {
		panic("Login failed")
	}

	// Protected request 1 - cookies are automatically included
	usersResp, err := session.Get("/users").Execute()
	
	// Protected request 2 - same session, same cookies
	profileResp, err := session.Get("/profile").Execute()
	
	fmt.Println(usersResp.JSON())
	fmt.Println(profileResp.JSON())
}
```

## Custom Cookie Management

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

type Client struct {
	session *fetch.Fetch
}

func NewClient(baseURL string) *Client {
	return &Client{
		session: fetch.Session().SetBaseURL(baseURL),
	}
}

func (c *Client) Login(username, password string) error {
	response, err := c.session.Post("/login", &fetch.Config{
		Body: map[string]interface{}{
			"username": username,
			"password": password,
		},
	}).Execute()
	
	if err != nil {
		return err
	}
	
	if !response.Ok() {
		return fmt.Errorf("login failed: %v", response.Error())
	}
	
	return nil
}

func (c *Client) GetUsers() (*fetch.Response, error) {
	return c.session.Get("/users").Execute()
}

func main() {
	client := NewClient("https://api.example.com")
	
	if err := client.Login("user", "pass"); err != nil {
		panic(err)
	}
	
	response, err := client.GetUsers()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.JSON())
}
```
