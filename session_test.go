package fetch

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestSession(t *testing.T) {
	f := Session()
	testify.Assert(t, f != nil, "Expected error, got nil")
}
