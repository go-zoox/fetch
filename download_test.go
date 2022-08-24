package fetch

import (
	"os"
	"testing"

	"github.com/go-zoox/testify"
)

func TestDownload(t *testing.T) {
	response, err := Download("https://httpbin.zcorky.com/image", "/tmp/image.webp")
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected status code 200, got", response.Status)
	}

	stat, err := os.Stat("/tmp/image.webp")
	if err != nil {
		t.Error(err)
	}

	if stat.Size() == 0 {
		t.Error("Expected file size not 0, got 0")
	}
}

func TestDownloadParamsError(t *testing.T) {
	_, err := Download("", "/tmp/image.webp")
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Download("https://httpbin.zcorky.com/image", "/tmp/image.webp", &Config{}, &Config{})
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Download("https://httpbin.zcorky.com/image", "/tmp/image.webp", &Config{})
	testify.Assert(t, err == nil, "Expected nil, got error")
}
