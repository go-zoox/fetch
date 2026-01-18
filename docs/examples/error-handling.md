# Error Handling Examples

## Basic Error Handling

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/status/404")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	if !response.Ok() {
		fmt.Println("Response error:", response.Error())
		return
	}

	fmt.Println(response.JSON())
}
```

## Check Response Status

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

switch response.StatusCode() {
case 200:
	fmt.Println("Success:", response.JSON())
case 404:
	fmt.Println("User not found")
case 500:
	fmt.Println("Server error:", response.Error())
default:
	fmt.Printf("Unexpected status: %d\n", response.StatusCode())
}
```

## Detailed Error Handling

```go
func fetchUser(id string) (map[string]interface{}, error) {
	response, err := fetch.Get(fmt.Sprintf("https://api.example.com/users/%s", id))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if !response.Ok() {
		return nil, fmt.Errorf("request failed with status %d: %s", 
			response.StatusCode(), response.String())
	}

	var user map[string]interface{}
	if err := response.UnmarshalJSON(&user); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return user, nil
}
```

## Handling Network Errors

```go
func safeFetch(url string) (*fetch.Response, error) {
	response, err := fetch.Get(url)
	if err != nil {
		// Handle network errors, timeouts, etc.
		return nil, fmt.Errorf("network error: %w", err)
	}

	if !response.Ok() {
		// Handle HTTP errors
		return nil, fmt.Errorf("HTTP error %d: %s", 
			response.StatusCode(), response.String())
	}

	return response, nil
}
```

## Error Handling with Retry

```go
func fetchWithRetry(url string, maxRetries int) (*fetch.Response, error) {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		response, err := fetch.Get(url)
		if err != nil {
			lastErr = err
			continue
		}
		
		if response.Ok() {
			return response, nil
		}
		
		// Retry only on 5xx errors
		if response.StatusCode() >= 500 {
			lastErr = response.Error()
			continue
		}
		
		// Don't retry on 4xx errors
		return nil, response.Error()
	}
	
	return nil, fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}
```

## Error Handling for Different Status Codes

```go
func handleResponse(response *fetch.Response) error {
	if response.Ok() {
		return nil
	}

	switch response.StatusCode() {
	case 400:
		return fmt.Errorf("bad request: %s", response.String())
	case 401:
		return fmt.Errorf("unauthorized - please check your credentials")
	case 403:
		return fmt.Errorf("forbidden - you don't have permission")
	case 404:
		return fmt.Errorf("resource not found")
	case 429:
		return fmt.Errorf("rate limited - too many requests")
	case 500, 502, 503:
		return fmt.Errorf("server error - please try again later")
	default:
		return fmt.Errorf("unexpected error %d: %s", 
			response.StatusCode(), response.String())
	}
}

func main() {
	response, err := fetch.Get("https://api.example.com/users/1")
	if err != nil {
		panic(err)
	}

	if err := handleResponse(response); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response.JSON())
}
```
