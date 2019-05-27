package apptweak

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
)

func setContentType(f string) string {
	ext := filepath.Ext(f)
	switch ext {
	case ".json":
		return "application/json"
	case ".html":
		return "text/plain"
	default:
		panic(fmt.Errorf("Unexpected or no file extension: %v", ext))
	}
}

type RewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// note that url.URL.ResolveReference doesn't work here
	// since t.u is an absolute url
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}
