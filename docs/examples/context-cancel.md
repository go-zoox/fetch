# Context Cancellation Examples

## Cancel Request

Cancel a request using context:

```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/go-zoox/fetch"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	f := fetch.New()
	f.SetBaseURL("https://httpbin.zcorky.com")
	f.SetURL("/delay/10")
	f.SetContext(ctx)

	// Start request in goroutine
	done := make(chan bool)
	go func() {
		response, err := f.Execute()
		if err != nil {
			fmt.Println("Request cancelled or failed:", err)
		} else {
			fmt.Println(response.JSON())
		}
		done <- true
	}()

	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("Request cancelled")
	
	<-done
}
```

## Context with Timeout

Use context timeout to automatically cancel long-running requests:

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
	f.SetURL("https://httpbin.zcorky.com/delay/10")

	response, err := f.Execute()
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else {
			fmt.Println("Request failed:", err)
		}
		return
	}

	fmt.Println(response.JSON())
}
```

## Cancelling Multiple Requests

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"github.com/go-zoox/fetch"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	urls := []string{
		"https://httpbin.zcorky.com/delay/5",
		"https://httpbin.zcorky.com/delay/5",
		"https://httpbin.zcorky.com/delay/5",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			
			f := fetch.New()
			f.SetContext(ctx)
			f.SetURL(u)
			
			response, err := f.Execute()
			if err != nil {
				fmt.Printf("Request to %s failed: %v\n", u, err)
			} else {
				fmt.Printf("Request to %s succeeded\n", u)
			}
		}(url)
	}

	// Cancel all requests after 2 seconds
	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("All requests cancelled")

	wg.Wait()
}
```

## Graceful Shutdown with Context

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-zoox/fetch"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		<-sigChan
		fmt.Println("Shutting down...")
		cancel()
	}()

	f := fetch.New()
	f.SetContext(ctx)
	f.SetURL("https://httpbin.zcorky.com/delay/60")

	response, err := f.Execute()
	if err != nil {
		if err == context.Canceled {
			fmt.Println("Request was cancelled during shutdown")
		} else {
			fmt.Println("Request failed:", err)
		}
		return
	}

	fmt.Println(response.JSON())
}
```
