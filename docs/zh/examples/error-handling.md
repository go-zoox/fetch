# 错误处理示例

## 基本错误处理

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/status/404")
	if err != nil {
		fmt.Println("请求错误:", err)
		return
	}

	if !response.Ok() {
		fmt.Println("响应错误:", response.Error())
		return
	}

	fmt.Println(response.JSON())
}
```

## 检查响应状态

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

switch response.StatusCode() {
case 200:
	fmt.Println("成功:", response.JSON())
case 404:
	fmt.Println("用户未找到")
case 500:
	fmt.Println("服务器错误:", response.Error())
default:
	fmt.Printf("意外状态: %d\n", response.StatusCode())
}
```

## 详细错误处理

```go
func fetchUser(id string) (map[string]interface{}, error) {
	response, err := fetch.Get(fmt.Sprintf("https://api.example.com/users/%s", id))
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	if !response.Ok() {
		return nil, fmt.Errorf("请求失败，状态码 %d: %s", 
			response.StatusCode(), response.String())
	}

	var user map[string]interface{}
	if err := response.UnmarshalJSON(&user); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return user, nil
}
```

## 处理网络错误

```go
func safeFetch(url string) (*fetch.Response, error) {
	response, err := fetch.Get(url)
	if err != nil {
		// 处理网络错误、超时等
		return nil, fmt.Errorf("网络错误: %w", err)
	}

	if !response.Ok() {
		// 处理 HTTP 错误
		return nil, fmt.Errorf("HTTP 错误 %d: %s", 
			response.StatusCode(), response.String())
	}

	return response, nil
}
```
