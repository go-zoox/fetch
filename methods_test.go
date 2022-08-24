package fetch

import "testing"

func TestMethodStream(t *testing.T) {
	f := New()
	f.Stream("http://example.com")
}
