# Timeout & Retry

## Timeout

Set a timeout for requests to prevent hanging:

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Timeout: 5 * time.Second,
})
```

### Using Fetch Instance

```go
f := fetch.New()
f.SetTimeout(10 * time.Second)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

### Context with Timeout

You can also use Go's context for more control:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

f := fetch.New()
f.SetContext(ctx)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

## Retry

Use the `Retry` method to retry failed requests:

```go
f := fetch.New()
f.SetURL("https://httpbin.zcorky.com/get")

response, err := f.Retry(func(nf *fetch.Fetch) {
	// Modify the request before retry if needed
	nf.SetHeader("X-Retry-Count", "1")
})
```

### Custom Retry Logic

For more complex retry logic, you can implement your own:

```go
func fetchWithRetry(url string, maxRetries int) (*fetch.Response, error) {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		response, err := fetch.Get(url)
		if err == nil && response.Ok() {
			return response, nil
		}
		
		lastErr = err
		time.Sleep(time.Second * time.Duration(i+1)) // Exponential backoff
	}
	
	return nil, lastErr
}
```

## Best Practices

1. **Always set timeouts** to prevent hanging requests
2. **Use context cancellation** for more fine-grained control
3. **Implement retry with backoff** for transient failures
4. **Log retry attempts** for debugging
