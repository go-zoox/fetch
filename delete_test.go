package fetch

import (
	"fmt"
	"testing"
)

func Test_Delete(t *testing.T) {
	response, err := Delete("https://httpbin.zcorky.com/Delete", &Config{
		Body: map[string]interface{}{
			"foo":     "bar",
			"foo2":    "bar2",
			"number":  1,
			"boolean": true,
			"array": []string{
				"foo3",
				"bar3",
			},
			"nest": map[string]string{
				"foo4": "bar4",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected status code 200, got", response.Status)
	}

	if response.Headers.Get("content-type") != "application/json; charset=utf-8" {
		t.Error("Expected content-type application/json; charset=utf-8, got", response.Headers.Get("content-type"))
	}

	if response.Headers.Get("server") != "openresty" {
		t.Error("Expected server openresty, got", response.Headers.Get("server"))
	}

	if response.Get("url").String() != "/Delete" {
		t.Error("Expected url /Delete, got", response.Get("url").String())
	}

	if response.Get("method").String() != "DELETE" {
		t.Error("Expected method DELETE, got", response.Get("method").String())
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

	if response.Get("body.foo").String() != "bar" {
		t.Error("Expected body.foo bar, got", response.Get("body.foo").String())
	}

	if response.Get("body.foo2").String() != "bar2" {
		t.Error("Expected body.foo2 bar2, got", response.Get("body.foo2").String())
	}

	if response.Get("body.number").Int() != 1 {
		t.Error("Expected body.number 1, got", response.Get("body.number").String())
	}

	if response.Get("body.boolean").Bool() != true {
		t.Error("Expected body.boolean true, got", response.Get("body.boolean").String())
	}

	if response.Get("body.array.0").String() != "foo3" {
		t.Error("Expected body.array foo3, got", response.Get("body.array").String())
	}

	if response.Get("body.array.1").String() != "bar3" {
		t.Error("Expected body.array bar3, got", response.Get("body.array").String())
	}

	if response.Get("body.nest.foo4").String() != "bar4" {
		t.Error("Expected body.nest.foo4 bar4, got", response.Get("body.nest.foo4").String())
	}
}

func Test_Delete_With_Header(t *testing.T) {
	response, err := Delete("https://httpbin.zcorky.com/Delete", &Config{
		Headers: map[string]string{
			"X-CUSTOM-VAR":   "custom-value",
			"x-custom-var-2": "custom-value-2",
		},
	})
	if err != nil {
		t.Error(err)
	}

	if response.Get("headers.x-custom-var").String() != "custom-value" {
		t.Error("Expected x-custom-var custom-value, got", response.Get("headers.x-custom-var").String())
	}

	if response.Get("headers.x-custom-var-2").String() != "custom-value-2" {
		t.Error("Expected x-custom-var-2 custom-value, got", response.Get("headers.x-custom-var").String())
	}
}

func Test_Delete_With_Query(t *testing.T) {
	response, err := Delete("https://httpbin.zcorky.com/Delete", &Config{
		Query: map[string]string{
			"foo":  "bar",
			"foo2": "bar2",
		},
	})
	if err != nil {
		t.Error(err)
	}

	if response.Get("query.foo").String() != "bar" {
		t.Error("Expected foo bar, got", response.Get("query.foo").String())
	}

	if response.Get("query.foo2").String() != "bar2" {
		t.Error("Expected foo2 bar2, got", response.Get("query.foo2").String())
	}
}
