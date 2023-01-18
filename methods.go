package fetch

import (
	"io"

	"github.com/go-zoox/headers"
)

// Head is http.head
func (f *Fetch) Head(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(HEAD).
		SetURL(url)
}

// Get is http.get
func (f *Fetch) Get(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(GET).
		SetURL(url)
}

// Post is http.post
func (f *Fetch) Post(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(POST).
		SetURL(url)
}

// Put is http.put
func (f *Fetch) Put(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(PUT).
		SetURL(url)
}

// Patch is http.patch
func (f *Fetch) Patch(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(PATCH).
		SetURL(url)
}

// Delete is http.delete
func (f *Fetch) Delete(url string, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(DELETE).
		SetURL(url)
}

// Download downloads file by url
func (f *Fetch) Download(url string, filepath string, config ...*Config) *Fetch {
	return f.
		SetHeader(headers.AcceptEncoding, "gzip").
		SetConfig(config...).
		SetMethod(GET).
		SetURL(url).
		SetDownloadFilePath(filepath)
}

// Upload upload a file
func (f *Fetch) Upload(url string, file io.Reader, config ...*Config) *Fetch {
	return f.
		SetConfig(config...).
		SetMethod(POST).
		SetURL(url).
		SetHeader(headers.ContentType, "multipart/form-data").
		SetBody(map[string]interface{}{
			"file": file,
		})
}

// Stream ...
func (f *Fetch) Stream(url string, config ...*Config) *Fetch {
	var cfg *Config = &Config{}
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Method == "" {
		cfg.Method = GET
	}

	cfg.IsStream = true

	return f.
		SetConfig(cfg).
		SetURL(url)
}

// func (f *Fetch) JSON() *Response {
// 	f.SetHeader("accept", "application/json")
// 	return f.Execute()
// }
