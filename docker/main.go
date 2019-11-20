package main

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"

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
	var isBatch bool

	switch r.Method {
	case http.MethodGet:
		isBatch = false
	case http.MethodPost:
		isBatch = true
	default:
		http.Error(rw, "Unexpected Method", http.StatusBadRequest)
		return
	}

	switch {
	case r.URL.Path == n1urlPath:
		if isBatch {
			batch(rw, r, bsn.FirstNormalizeURL)
		} else {
			normalize(rw, r, bsn.FirstNormalizeURL)
		}
	case r.URL.Path == n2urlPath:
		if isBatch {
			batch(rw, r, bsn.SecondNormalizeURL)
		} else {
			normalize(rw, r, bsn.SecondNormalizeURL)
		}
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

func normalize(rw http.ResponseWriter, r *http.Request, f func(u *url.URL) string) {
	u, err := getURL(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	data := []byte(f(u))
	if _, err := rw.Write(data); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

func batch(rw http.ResponseWriter, r *http.Request, f func(u *url.URL) string) {
	buf := &bytes.Buffer{}

	s := bufio.NewScanner(r.Body)
	defer r.Body.Close()

	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		u, err := url.Parse(s.Text())
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		buf.WriteString(f(u) + "\n")
	}

	if s.Err() != nil {
		http.Error(rw, s.Err().Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(rw, buf); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
