package fetch

import (
	"testing"

	"github.com/go-zoox/testify"
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

func TestHeadParamsError(t *testing.T) {
	_, err := Head("")
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Head("https://httpbin.zcorky.com/image", &Config{}, &Config{})
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Head("https://httpbin.zcorky.com/image", &Config{})
	testify.Assert(t, err == nil, "Expected nil, got error")
}
