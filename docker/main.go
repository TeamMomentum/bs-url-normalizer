package main

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	bsn "github.com/TeamMomentum/bs-url-normalizer/lib/urls"
)

const (
	n1urlPath = "/n1url"
	n2urlPath = "/n2url"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := &handler{}

	if err := http.ListenAndServe(":"+port, h); err != nil {
		panic(err)
	}
}

type handler struct{}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch {
	case strings.HasPrefix(r.URL.Path, n1urlPath):
		u, err := getURL(r)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ret := []byte(bsn.FirstNormalizeURL(u))
		if _, err := rw.Write(ret); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
	case strings.HasPrefix(r.URL.Path, n2urlPath):
		u, err := getURL(r)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ret := []byte(bsn.SecondNormalizeURL(u))
		if _, err := rw.Write(ret); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
	case strings.HasPrefix(r.URL.Path, n2urlPath):
	default:
		http.Error(rw, "404", http.StatusNotFound)
	}
}

func getURL(r *http.Request) (*url.URL, error) {
	qurl := r.URL.Query().Get("url")
	src, err := url.QueryUnescape(qurl)
	if err != nil {
		return nil, err
	}

	return url.Parse(src)
}
