# Response Handling

Fetch provides a convenient `Response` object for handling HTTP responses.

## Basic Usage

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")
if err != nil {
	panic(err)
}

// Get status code
status := response.Status

// Get response body as string
body := response.String()

// Get response body as JSON
json, err := response.JSON()
```

## JSON Parsing

Fetch uses [gjson](https://github.com/tidwall/gjson) for JSON parsing, allowing you to access nested values without unmarshaling.

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")

// Access JSON values using gjson syntax
url := response.Get("url")
method := response.Get("method")

// Access nested values
value := response.Get("data.user.name")
```

## Unmarshal to Struct

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

response, err := fetch.Get("https://api.example.com/users/1")

var user User
err = response.UnmarshalJSON(&user)
```

## Response Methods

### Status Check

```go
if response.Ok() {
	// Status code is 2xx
}
```

### Get Status Information

```go
statusCode := response.StatusCode()
statusText := response.StatusText()
```

### Get Headers

```go
contentType := response.ContentType()
location := response.Location()
contentLength := response.ContentLength()
```

### Get Response Error

```go
if !response.Ok() {
	err := response.Error()
	// err contains status code and response body
}
```

## Example

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	var user User
	if err := response.UnmarshalJSON(&user); err != nil {
		panic(err)
	}
	fmt.Printf("User: %s\n", user.Name)
} else {
	fmt.Printf("Error: %v\n", response.Error())
}
```
