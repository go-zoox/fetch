# 示例

本节提供了常见用例的实用示例。

## 概述

- [基础用法](./basic) - 简单的 GET 和 POST 请求
- [HTTP 方法](./http-methods) - 所有 HTTP 方法（GET、POST、PUT、PATCH、DELETE、HEAD）
- [认证](./auth) - Basic 认证、Bearer Token、自定义请求头
- [文件操作](./file-operations) - 上传和下载文件
- [超时与重试](./timeout-retry) - 设置超时和重试失败的请求
- [代理](./proxy) - 使用 HTTP、HTTPS 和 SOCKS5 代理
- [流式传输](./stream) - 流式传输响应数据
- [会话与 Cookie](./session-cookies) - 会话管理和 Cookie 处理
- [Context 取消](./context-cancel) - 使用 context 取消请求
- [错误处理](./error-handling) - 处理错误和响应状态码

## 快速示例

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/get")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```
