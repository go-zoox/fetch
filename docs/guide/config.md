# Configuration

The `Config` struct provides various options to customize HTTP requests.

## Basic Configuration

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Timeout: 10 * time.Second,
	Headers: map[string]string{
		"User-Agent": "MyApp/1.0",
	},
})
```

## Configuration Options

### URL and Method

- `URL`: The target URL for the request
- `Method`: HTTP method (GET, POST, PUT, PATCH, DELETE, HEAD)
- `BaseURL`: Base URL for relative paths

### Headers

```go
&fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer token",
		"Content-Type": "application/json",
	},
}
```

### Query Parameters

```go
&fetch.Config{
	Query: map[string]string{
		"page": "1",
		"limit": "10",
	},
}
```

### Request Body

```go
&fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"age": 30,
	},
}
```

### Timeout

```go
&fetch.Config{
	Timeout: 5 * time.Second,
}
```

### Proxy

```go
&fetch.Config{
	Proxy: "http://127.0.0.1:8080",
}
```

### Authentication

```go
// Basic Auth
&fetch.Config{
	BasicAuth: &fetch.BasicAuth{
		Username: "user",
		Password: "pass",
	},
}

// Or use Bearer Token via Headers
&fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer token",
	},
}
```

### TLS Configuration

```go
&fetch.Config{
	TLSCaCertFile: "/path/to/ca.crt",
	TLSInsecureSkipVerify: false,
}
```

### Unix Domain Socket

```go
&fetch.Config{
	UnixDomainSocket: "/var/run/docker.sock",
}
```

## Using Fetch Instance

You can also configure a Fetch instance and reuse it:

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")
f.SetTimeout(10 * time.Second)
f.SetHeader("Authorization", "Bearer token")

response, err := f.Get("/users").Execute()
```
