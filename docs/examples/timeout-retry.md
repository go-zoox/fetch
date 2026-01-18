# Timeout & Retry Examples

## Timeout

Set a timeout for requests to prevent hanging:

```go
package main

import (
	"fmt"
	"time"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## Timeout with Fetch Instance

```go
f := fetch.New()
f.SetTimeout(10 * time.Second)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

## Context with Timeout

Use Go's context for more control:

```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/go-zoox/fetch"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	f := fetch.New()
	f.SetContext(ctx)

	response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Println(response.JSON())
}
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

## Custom Retry Logic with Exponential Backoff

```go
package main

import (
	"fmt"
	"time"
	"github.com/go-zoox/fetch"
)

func fetchWithRetry(url string, maxRetries int) (*fetch.Response, error) {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		response, err := fetch.Get(url)
		if err == nil && response.Ok() {
			return response, nil
		}
		
		lastErr = err
		if i < maxRetries-1 {
			// Exponential backoff: 1s, 2s, 4s...
			backoff := time.Second * time.Duration(1<<uint(i))
			time.Sleep(backoff)
		}
	}
	
	return nil, fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

func main() {
	response, err := fetchWithRetry("https://httpbin.zcorky.com/get", 3)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.JSON())
}
```

## Retry with Different Strategies

```go
func retryWithExponentialBackoff(fn func() (*fetch.Response, error), maxRetries int) (*fetch.Response, error) {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		response, err := fn()
		if err == nil && response.Ok() {
			return response, nil
		}
		
		lastErr = err
		if i < maxRetries-1 {
			backoff := time.Second * time.Duration(1<<uint(i))
			time.Sleep(backoff)
		}
	}
	
	return nil, lastErr
}

func main() {
	response, err := retryWithExponentialBackoff(func() (*fetch.Response, error) {
		return fetch.Get("https://httpbin.zcorky.com/get")
	}, 3)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.JSON())
}
```
