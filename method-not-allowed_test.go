package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
	"sethwklein.net/go/errutil"
)

func Uses405(address string) (err error) {
	res, err := http.Get(address)
	defer errutil.AppendCall(&err, res.Body.Close)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return errors.New("server claims to support GET /")
	case 404:
		return errors.New("wat")
	case 405:
		return nil
	default:
		return fmt.Errorf("odd status code: %v", res.StatusCode)
	}
}

func MethodNotAllowedStandardLibrary() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "use POST", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("test body"))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close() // no return value

	return Uses405(ts.URL)
}

func TestMethodNotAllowedStandardLibrary(t *testing.T) {
	err := MethodNotAllowedStandardLibrary()
	if err != nil {
		t.Error(err)
	}
}

func MethodNotAllowedGoji() (err error) {
	mux := web.New()
	mux.Post("/", func(_ web.C, w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("test body"))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close() // no return value

	return Uses405(ts.URL)
}

func TestMethodNotAllowedGoji(t *testing.T) {
	err := MethodNotAllowedGoji()
	if err != nil {
		t.Error(err)
	}
}
