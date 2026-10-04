package fetch_test

import (
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

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

func TestRequestTLSOptionsDoNotChangeDefaultTransport(t *testing.T) {
	for _, option := range []string{"insecure verification", "certificate authority"} {
		t.Run(option, func(t *testing.T) {
			transport := isolatedDefaultTransport(t)
			roots := x509.NewCertPool()
			transport.TLSClientConfig = &tls.Config{RootCAs: roots}

			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("ok"))
			}))
			server.Config.ErrorLog = log.New(io.Discard, "", 0)
			server.StartTLS()
			defer server.Close()

			config := &fetch.Config{}
			if option == "insecure verification" {
				config.TLSInsecureSkipVerify = true
			} else {
				config.TLSCaCert = pem.EncodeToMemory(&pem.Block{
					Type:  "CERTIFICATE",
					Bytes: server.Certificate().Raw,
				})
			}

			response, err := fetch.Get(server.URL, config)
			if err != nil {
				t.Fatalf("Request with TLS option failed: %v", err)
			}
			if string(response.Body) != "ok" {
				t.Fatalf("Expected response body 'ok', got: %s", response.Body)
			}

			if transport.TLSClientConfig.InsecureSkipVerify {
				t.Error("Request disabled certificate verification on the default transport")
			}
			if transport.TLSClientConfig.RootCAs != roots {
				t.Error("Request replaced the default transport's certificate authorities")
			}
			if _, err := fetch.Get(server.URL); err == nil || !strings.Contains(err.Error(), "certificate signed by unknown authority") {
				t.Errorf("Expected subsequent request to reject the server certificate, got: %v", err)
			}
		})
	}
}

func TestRequestClientCertificateDoesNotChangeDefaultTransport(t *testing.T) {
	transport := isolatedDefaultTransport(t)
	defaultCert, _, _ := clientCertificate(t, "default-client")
	_, requestCertPEM, requestKeyPEM := clientCertificate(t, "request-client")

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.TLS.PeerCertificates[0].Subject.CommonName))
	}))
	server.TLS = &tls.Config{ClientAuth: tls.RequireAnyClientCert}
	server.StartTLS()
	defer server.Close()

	roots := x509.NewCertPool()
	roots.AddCert(server.Certificate())
	transport.TLSClientConfig = &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{defaultCert},
	}

	response, err := fetch.Get(server.URL, &fetch.Config{TLSCert: requestCertPEM, TLSKey: requestKeyPEM})
	if err != nil {
		t.Fatalf("Request with client certificate failed: %v", err)
	}
	if string(response.Body) != "request-client" {
		t.Errorf("Expected request client certificate, got: %s", response.Body)
	}
	if transport.TLSClientConfig.RootCAs != roots {
		t.Error("Request replaced the default transport's certificate authorities")
	}

	response, err = fetch.Get(server.URL)
	if err != nil {
		t.Fatalf("Subsequent request with default client certificate failed: %v", err)
	}
	if string(response.Body) != "default-client" {
		t.Errorf("Expected default client certificate on subsequent request, got: %s", response.Body)
	}
}

func TestRequestUnixSocketDoesNotChangeDefaultTransport(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unix socket test requires Unix")
	}
	isolatedDefaultTransport(t)

	socketDir, err := os.MkdirTemp("", "fetch-")
	if err != nil {
		t.Fatalf("Failed to create socket directory: %v", err)
	}
	defer os.RemoveAll(socketDir)
	socketPath := filepath.Join(socketDir, "http.sock")
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		t.Fatalf("Failed to listen on Unix socket: %v", err)
	}
	unixServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("unix"))
	}))
	unixServer.Listener.Close()
	unixServer.Listener = listener
	unixServer.Start()
	defer unixServer.Close()

	tcpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("tcp"))
	}))
	defer tcpServer.Close()

	response, err := fetch.Get(tcpServer.URL, &fetch.Config{UnixDomainSocket: "unix://" + socketPath})
	if err != nil {
		t.Fatalf("Request over Unix socket failed: %v", err)
	}
	if string(response.Body) != "unix" {
		t.Errorf("Expected Unix socket response, got: %s", response.Body)
	}

	response, err = fetch.Get(tcpServer.URL)
	if err != nil {
		t.Fatalf("Subsequent TCP request failed: %v", err)
	}
	if string(response.Body) != "tcp" {
		t.Errorf("Expected TCP response on subsequent request, got: %s", response.Body)
	}

	t.Run("proxy with custom default transport", func(t *testing.T) {
		original := http.DefaultTransport
		http.DefaultTransport = &struct{ http.RoundTripper }{original}
		defer func() { http.DefaultTransport = original }()

		response, err := fetch.Get(tcpServer.URL, &fetch.Config{
			Proxy:            "http://proxy.invalid:8080",
			UnixDomainSocket: socketPath,
		})
		if err != nil {
			t.Fatalf("Request with proxy and custom default transport failed: %v", err)
		}
		if string(response.Body) != "unix" {
			t.Errorf("Expected Unix socket response with proxy, got: %s", response.Body)
		}
	})
}

func TestRequestPrivateTransportClosesConnection(t *testing.T) {
	for _, test := range []struct {
		name   string
		http2  bool
		stream bool
	}{
		{name: "HTTP1 buffered"},
		{name: "HTTP1 streamed", stream: true},
		{name: "HTTP2 buffered", http2: true},
		{name: "HTTP2 streamed", http2: true, stream: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			transport := isolatedDefaultTransport(t)
			transport.DisableKeepAlives = false
			transport.IdleConnTimeout = 0
			transport.ForceAttemptHTTP2 = test.http2

			closed := make(chan struct{})
			release := make(chan struct{})
			var closedOnce, releaseOnce sync.Once
			releaseBody := func() { releaseOnce.Do(func() { close(release) }) }
			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Test-Protocol", r.Proto)
				w.Write([]byte("prefix"))
				if test.stream {
					w.(http.Flusher).Flush()
					<-release
				}
				w.Write([]byte("suffix"))
			}))
			server.EnableHTTP2 = test.http2
			server.Config.ConnState = func(_ net.Conn, state http.ConnState) {
				if state == http.StateClosed {
					closedOnce.Do(func() { close(closed) })
				}
			}
			server.StartTLS()
			defer server.Close()
			defer releaseBody()

			response, err := fetch.Get(server.URL, &fetch.Config{
				TLSInsecureSkipVerify: true,
				IsStream:              test.stream,
			})
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			protocol := "HTTP/1.1"
			if test.http2 {
				protocol = "HTTP/2.0"
			}
			if got := response.Headers.Get("X-Test-Protocol"); got != protocol {
				t.Fatalf("Expected protocol %s, got: %s", protocol, got)
			}
			body := response.Body
			if test.stream {
				defer response.Stream.Close()
				body = make([]byte, len("prefix"))
				if _, err := io.ReadFull(response.Stream, body); err != nil {
					t.Fatalf("Failed to read stream after Execute returned: %v", err)
				}
				releaseBody()
				rest, err := io.ReadAll(response.Stream)
				if err != nil {
					t.Fatalf("Failed to finish stream: %v", err)
				}
				body = append(body, rest...)
				if err := response.Stream.Close(); err != nil {
					t.Fatalf("Failed to close stream: %v", err)
				}
			}
			if string(body) != "prefixsuffix" {
				t.Errorf("Expected complete response body, got: %s", body)
			}
			select {
			case <-closed:
			case <-time.After(2 * time.Second):
				t.Error("Private transport kept its connection open after the response was consumed")
			}
		})
	}
}

func isolatedDefaultTransport(t *testing.T) *http.Transport {
	t.Helper()
	original := http.DefaultTransport
	transport := original.(*http.Transport).Clone()
	transport.Proxy = nil
	transport.DisableKeepAlives = true
	http.DefaultTransport = transport
	t.Cleanup(func() {
		transport.CloseIdleConnections()
		http.DefaultTransport = original
	})
	return transport
}

func clientCertificate(t *testing.T, name string) (tls.Certificate, []byte, []byte) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate client key: %v", err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: name},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("Failed to create client certificate: %v", err)
	}
	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatalf("Failed to encode client key: %v", err)
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatalf("Failed to load client certificate: %v", err)
	}
	return cert, certPEM, keyPEM
}
