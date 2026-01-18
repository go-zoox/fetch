# Basic Usage Examples

## Simple GET Request

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

## PUT Request

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
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
response, err := fetch.Delete("https://httpbin.zcorky.com/delete")
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

## Handling Response

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	// Access JSON values
	name := response.Get("name")
	email := response.Get("email")
	
	// Or unmarshal to struct
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	var user User
	err := response.UnmarshalJSON(&user)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("User: %s (%s)\n", user.Name, user.Email)
} else {
	fmt.Printf("Error: %v\n", response.Error())
}
```
