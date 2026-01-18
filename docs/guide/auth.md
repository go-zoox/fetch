# Authentication

Fetch supports multiple authentication methods.

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

## Bearer Token

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

### Using SetAuthorization

```go
f := fetch.New()
f.SetAuthorization("Bearer your-token-here")

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

## Example: API with Authentication

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	f := fetch.New()
	f.SetBaseURL("https://api.example.com")
	f.SetBearerToken("your-api-token")

	response, err := f.Get("/users").Execute()
	if err != nil {
		panic(err)
	}

	if response.Ok() {
		fmt.Println(response.JSON())
	} else {
		fmt.Printf("Error: %v\n", response.Error())
	}
}
```
