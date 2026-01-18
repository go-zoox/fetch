# Context 取消示例

## 取消请求

使用 context 取消请求：

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

	// 在 goroutine 中启动请求
	done := make(chan bool)
	go func() {
		response, err := f.Execute()
		if err != nil {
			fmt.Println("请求已取消或失败:", err)
		} else {
			fmt.Println(response.JSON())
		}
		done <- true
	}()

	// 1 秒后取消
	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("请求已取消")
	
	<-done
}
```

## Context 超时

使用 context 超时自动取消长时间运行的请求：

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
			fmt.Println("请求超时")
		} else {
			fmt.Println("请求失败:", err)
		}
		return
	}

	fmt.Println(response.JSON())
}
```
