# HTTP Methods

Fetch supports all standard HTTP methods: GET, POST, PUT, PATCH, DELETE, and HEAD.

## GET

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")
```

With query parameters:

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Query: map[string]string{
		"page": "1",
	},
})
```

## POST

```go
response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PUT

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
		"name": "Jane",
		"email": "jane@example.com",
	},
})
```

## PATCH

```go
response, err := fetch.Patch("https://httpbin.zcorky.com/patch", &fetch.Config{
	Body: map[string]interface{}{
		"email": "newemail@example.com",
	},
})
```

## DELETE

```go
response, err := fetch.Delete("https://httpbin.zcorky.com/delete")
```

## HEAD

```go
response, err := fetch.Head("https://httpbin.zcorky.com/get")
```

## Using Fetch Instance

You can also use a Fetch instance for method chaining:

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
