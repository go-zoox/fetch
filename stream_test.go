package fetch

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestStream(t *testing.T) {
	_, err := Stream("")
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Stream("https://httpbin.zcorky.com/image", &Config{}, &Config{})
	testify.Assert(t, err != nil, "Expected error, got nil")

	_, err = Stream("https://httpbin.zcorky.com/image", &Config{})
	testify.Assert(t, err == nil, "Expected nil, got error")
}
