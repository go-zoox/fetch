# HTTP 方法

Fetch 支持所有标准 HTTP 方法：GET、POST、PUT、PATCH、DELETE 和 HEAD。

## GET

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get")
```

带查询参数：

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Query: map[string]string{
		"page": "1",
	},
})
```

## POST

```go
response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"email": "john@example.com",
	},
})
```

## PUT

```go
response, err := fetch.Put("https://httpbin.zcorky.com/put", &fetch.Config{
	Body: map[string]interface{}{
		"name": "Jane",
		"email": "jane@example.com",
	},
})
```

## PATCH

```go
response, err := fetch.Patch("https://httpbin.zcorky.com/patch", &fetch.Config{
	Body: map[string]interface{}{
		"email": "newemail@example.com",
	},
})
```

## DELETE

```go
response, err := fetch.Delete("https://httpbin.zcorky.com/delete")
```

## HEAD

```go
response, err := fetch.Head("https://httpbin.zcorky.com/get")
```

## 使用 Fetch 实例

您也可以使用 Fetch 实例进行方法链式调用：

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")

// GET
response, err := f.Get("/users").Execute()

// POST
response, err := f.Post("/users", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
	},
}).Execute()
```
