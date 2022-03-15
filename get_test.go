package fetch

import (
	"fmt"
	"testing"
)

func Test_Get(t *testing.T) {
	response := Get("https://httpbin.zcorky.com/get")

	if response.Status != 200 {
		t.Error("Expected status code 200, got", response.Status)
	}

	if response.Headers.Get("content-type") != "application/json; charset=utf-8" {
		t.Error("Expected content-type application/json; charset=utf-8, got", response.Headers.Get("content-type"))
	}

	if response.Headers.Get("server") != "openresty" {
		t.Error("Expected server openresty, got", response.Headers.Get("server"))
	}

	if response.Get("url").String() != "/get" {
		t.Error("Expected url /get, got", response.Get("url").String())
	}

	if response.Get("method").String() != "GET" {
		t.Error("Expected method GET, got", response.Get("method").String())
	}

	if response.Get("headers.host").String() != "httpbin.zcorky.com" {
		t.Error("Expected Host httpbin.zcorky.com, got", response.Get("headers.Host").String())
	}

	if response.Get("headers.accept-encoding").String() != "gzip" {
		t.Error("Expected accept-encoding gzip, got", response.Get("headers.accept-encoding").String())
	}

	if response.Get("headers.user-agent").String() != DefaultUserAgent() {
		t.Error(fmt.Sprintf("Expected user-agent %s, got", DefaultUserAgent()), response.Get("headers.user-agent").String())
	}

	if response.Get("headers.connection").String() != "close" {
		t.Error("Expected connection close, got", response.Get("headers.connection").String())
	}

	if response.Get("origin").String() != "https://httpbin.zcorky.com" {
		t.Error("Expected origin https://httpbin.zcorky.com, got", response.Get("origin").String())
	}
}

func Test_Get_With_Header(t *testing.T) {
	response := Get("https://httpbin.zcorky.com/get", &Config{
		Headers: map[string]string{
			"X-CUSTOM-VAR":   "custom-value",
			"x-custom-var-2": "custom-value-2",
		},
	})

	// fmt.Println("raw: ", response.JSON())

	if response.Get("headers.x-custom-var").String() != "custom-value" {
		t.Error("Expected x-custom-var custom-value, got", response.Get("headers.x-custom-var").String())
	}

	if response.Get("headers.x-custom-var-2").String() != "custom-value-2" {
		t.Error("Expected x-custom-var-2 custom-value, got", response.Get("headers.x-custom-var").String())
	}
}

func Test_Get_With_Query(t *testing.T) {
	response := Get("https://httpbin.zcorky.com/get", &Config{
		Query: map[string]string{
			"foo":  "bar",
			"foo2": "bar2",
		},
	})

	if response.Get("query.foo").String() != "bar" {
		t.Error("Expected foo bar, got", response.Get("query.foo").String())
	}

	if response.Get("query.foo2").String() != "bar2" {
		t.Error("Expected foo2 bar2, got", response.Get("query.foo2").String())
	}
}
