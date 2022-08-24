package fetch

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestConfigMerge(t *testing.T) {
	// nil
	cfg := Config{}
	cfg.Merge(nil)

	// merge params
	testify.Equal(t, "", cfg.Headers.Get("key"))
	cfg.Merge(&Config{
		Params: Params{
			"key": "value",
		},
	})
	testify.Equal(t, "value", cfg.Params.Get("key"))

	// merge query
	testify.Equal(t, "", cfg.Query.Get("key"))
	cfg.Merge(&Config{
		Query: Query{
			"key": "value",
		},
	})
	testify.Equal(t, "value", cfg.Query.Get("key"))

	// merge headers
	testify.Equal(t, "", cfg.Headers.Get("key"))
	cfg.Merge(&Config{
		Headers: Headers{
			"key": "value",
		},
	})
	testify.Equal(t, "value", cfg.Headers.Get("key"))

	// merge base url
	testify.Equal(t, "", cfg.BaseURL)
	cfg.Merge(&Config{
		BaseURL: "http://example.com",
	})
	testify.Equal(t, "http://example.com", cfg.BaseURL)

	// merge timeout
	testify.Equal(t, 0, cfg.Timeout)
	cfg.Merge(&Config{
		Timeout: 1,
	})
	testify.Equal(t, 1, cfg.Timeout)

	// merge download file path
	testify.Equal(t, "", cfg.DownloadFilePath)
	cfg.Merge(&Config{
		DownloadFilePath: "path",
	})
	testify.Equal(t, "path", cfg.DownloadFilePath)

	// merge proxy
	testify.Equal(t, "", cfg.Proxy)
	cfg.Merge(&Config{
		Proxy: "http://example.com",
	})
	testify.Equal(t, "http://example.com", cfg.Proxy)

	// merge is stream
	testify.Equal(t, false, cfg.IsStream)
	cfg.Merge(&Config{
		IsStream: true,
	})
	testify.Equal(t, true, cfg.IsStream)
}

func TestConfigHeaders(t *testing.T) {
	headers := Headers{}
	testify.Equal(t, headers.Get("key"), "", "Expected empty string")

	headers.Set("key", "value")
	testify.Equal(t, headers.Get("key"), "value", "Expected value")
}

func TestConfigQuery(t *testing.T) {
	query := Query{}
	testify.Equal(t, query.Get("key"), "", "Expected empty string")

	query.Set("key", "value")
	testify.Equal(t, query.Get("key"), "value", "Expected value")
}

func TestConfigParams(t *testing.T) {
	params := Params{}
	testify.Equal(t, params.Get("key"), "", "Expected empty string")

	params.Set("key", "value")
	testify.Equal(t, params.Get("key"), "value", "Expected value")
}
