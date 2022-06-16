package fetch

import (
	"testing"
)

func TestHead(t *testing.T) {
	response, err := Head("https://httpbin.zcorky.com")
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected status code 200, got", response.Status)
	}

	if response.Headers.Get("content-type") != "text/plain; charset=utf-8" {
		t.Error("Expected content-type text/plain; charset=utf-8, got", response.Headers.Get("content-type"))
	}

	if response.Headers.Get("server") == "" {
		t.Error("Expected server not empty, got empty")
	}
}
