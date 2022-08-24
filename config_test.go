package fetch

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestHeaders(t *testing.T) {
	headers := Headers{}
	testify.Equal(t, headers.Get("key"), "", "Expected empty string")

	headers.Set("key", "value")
	testify.Equal(t, headers.Get("key"), "value", "Expected value")
}

func TestQuery(t *testing.T) {
	query := Query{}
	testify.Equal(t, query.Get("key"), "", "Expected empty string")

	query.Set("key", "value")
	testify.Equal(t, query.Get("key"), "value", "Expected value")
}

func TestParams(t *testing.T) {
	params := Params{}
	testify.Equal(t, params.Get("key"), "", "Expected empty string")

	params.Set("key", "value")
	testify.Equal(t, params.Get("key"), "value", "Expected value")
}
