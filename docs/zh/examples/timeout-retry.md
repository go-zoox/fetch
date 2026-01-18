# 超时与重试示例

## 超时

设置请求超时以防止请求挂起：

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

## 使用 Fetch 实例设置超时

```go
f := fetch.New()
f.SetTimeout(10 * time.Second)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

## 使用 Context 设置超时

使用 Go 的 context 获得更多控制：

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
		fmt.Println("请求失败:", err)
		return
	}

	fmt.Println(response.JSON())
}
```

## 重试

使用 `Retry` 方法重试失败的请求：

```go
f := fetch.New()
f.SetURL("https://httpbin.zcorky.com/get")

response, err := f.Retry(func(nf *fetch.Fetch) {
	// 如果需要，在重试前修改请求
	nf.SetHeader("X-Retry-Count", "1")
})
```

## 自定义重试逻辑（指数退避）

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
			// 指数退避：1s, 2s, 4s...
			backoff := time.Second * time.Duration(1<<uint(i))
			time.Sleep(backoff)
		}
	}
	
	return nil, fmt.Errorf("在 %d 次重试后失败: %w", maxRetries, lastErr)
}

func main() {
	response, err := fetchWithRetry("https://httpbin.zcorky.com/get", 3)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.JSON())
}
```
