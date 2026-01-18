# 响应处理

Fetch 提供了一个方便的 `Response` 对象来处理 HTTP 响应。

## 基本用法

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")
if err != nil {
	panic(err)
}

// 获取状态码
status := response.Status

// 获取响应体作为字符串
body := response.String()

// 获取响应体作为 JSON
json, err := response.JSON()
```

## JSON 解析

Fetch 使用 [gjson](https://github.com/tidwall/gjson) 进行 JSON 解析，允许您在不反序列化的情况下访问嵌套值。

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")

// 使用 gjson 语法访问 JSON 值
url := response.Get("url")
method := response.Get("method")

// 访问嵌套值
value := response.Get("data.user.name")
```

## 反序列化为结构体

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

response, err := fetch.Get("https://api.example.com/users/1")

var user User
err = response.UnmarshalJSON(&user)
```

## 响应方法

### 状态检查

```go
if response.Ok() {
	// 状态码为 2xx
}
```

### 获取状态信息

```go
statusCode := response.StatusCode()
statusText := response.StatusText()
```

### 获取请求头

```go
contentType := response.ContentType()
location := response.Location()
contentLength := response.ContentLength()
```

### 获取响应错误

```go
if !response.Ok() {
	err := response.Error()
	// err 包含状态码和响应体
}
```

## 示例

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	var user User
	if err := response.UnmarshalJSON(&user); err != nil {
		panic(err)
	}
	fmt.Printf("用户: %s\n", user.Name)
} else {
	fmt.Printf("错误: %v\n", response.Error())
}
```
