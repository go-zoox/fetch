# Advanced Features

## HTTP/2 Support

Enable HTTP/2:

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	HTTP2: true,
})
```

## TLS Configuration

### Custom CA Certificate

```go
response, err := fetch.Get("https://internal-api.example.com", &fetch.Config{
	TLSCaCertFile: "/path/to/ca.crt",
})
```

Or from bytes:

```go
caCert, _ := os.ReadFile("/path/to/ca.crt")

response, err := fetch.Get("https://internal-api.example.com", &fetch.Config{
	TLSCaCert: caCert,
})
```

### Client Certificate Authentication

```go
response, err := fetch.Get("https://api.example.com", &fetch.Config{
	TLSCertFile: "/path/to/client.crt",
	TLSKeyFile:  "/path/to/client.key",
})
```

### Skip TLS Verification (Not Recommended for Production)

```go
response, err := fetch.Get("https://self-signed.example.com", &fetch.Config{
	TLSInsecureSkipVerify: true,
})
```

## Unix Domain Socket

Connect via Unix domain socket:

```go
response, err := fetch.Get("http://localhost/info", &fetch.Config{
	UnixDomainSocket: "/var/run/docker.sock",
})
```

## Streaming

Stream response data:

```go
f := fetch.New()
f.SetURL("https://example.com/stream")
f.SetMethod("GET")

config, _ := f.Config()
config.IsStream = true

response, err := f.SetConfig(config).Execute()
if err != nil {
	panic(err)
}

defer response.Stream.Close()

// Read stream
buf := make([]byte, 1024)
for {
	n, err := response.Stream.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		panic(err)
	}
	
	// Process chunk
	processChunk(buf[:n])
}
```

## Context Cancellation

Cancel requests using context:

```go
ctx, cancel := context.WithCancel(context.Background())

f := fetch.New()
f.SetContext(ctx)
f.SetURL("https://slow-api.example.com")

// Start request in goroutine
go func() {
	response, err := f.Execute()
	if err != nil {
		fmt.Println("Request cancelled or failed:", err)
		return
	}
	fmt.Println(response.JSON())
}()

// Cancel after 1 second
time.Sleep(1 * time.Second)
cancel()
```

## Custom Headers

```go
f := fetch.New()
f.SetHeader("X-Custom-Header", "value")
f.SetHeader("User-Agent", "MyApp/1.0")
f.SetAccept("application/json")
f.SetContentType("application/json")
```

## Chaining Methods

Fetch supports method chaining for cleaner code:

```go
response, err := fetch.New().
	SetBaseURL("https://api.example.com").
	SetBearerToken("token").
	SetTimeout(10*time.Second).
	Get("/users").
	Execute()
```

## Session Management

Maintain cookies across requests:

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")

// Login
f.Post("/login", &fetch.Config{
	Body: map[string]interface{}{
		"username": "user",
		"password": "pass",
	},
}).Execute()

// Subsequent requests will use cookies from login
response, err := f.Get("/protected").Execute()
```
