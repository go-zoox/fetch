package fetch_test

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-zoox/fetch"
)

func TestRequestBodyCompression(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") != "gzip" {
			t.Fatalf("Expected Content-Encoding to be gzip, got: %s", r.Header.Get("Content-Encoding"))
		}

		gr, err := gzip.NewReader(r.Body)
		if err != nil {
			t.Fatalf("Failed to create gzip reader: %v", err)
		}
		defer gr.Close()

		body, err := io.ReadAll(gr)
		if err != nil {
			t.Fatalf("Failed to read compressed body: %v", err)
		}

		expected := `{"message":"hello"}`
		if string(body) != expected {
			t.Errorf("Decompressed body mismatch.\nExpected: %s\nGot: %s", expected, string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	res, err := fetch.Post(server.URL, &fetch.Config{
		Body: map[string]string{"message": "hello"},
		// Enable compression
		CompressRequest: true,
	})
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if res.Status != http.StatusOK {
		t.Errorf("Expected status 200 OK, got: %d", res.Status)
	}

	if strings.TrimSpace(string(res.Body)) != "ok" {
		t.Errorf("Expected response body 'ok', got: %s", res.Body)
	}
}
