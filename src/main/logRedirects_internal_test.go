package main

import (
	"testing"
	"net/http"
	"net/url"
)

func TestLogRedirects(t *testing.T) {
	var via = []*http.Request{
		&http.Request{
			URL: &url.URL{
				Host: "google.com",
			},  
		},
		&http.Request{
			URL: &url.URL{
				Host: "google.be",
			},
		},
	}

	t.Run("with RemoveQuery", func(t *testing.T) {
		RemoveQuery = true
		
		var req = http.Request{
			URL: &url.URL{
				RawQuery: "?id=1",
			},
		}

		err := logRedirect(&req, via)
		if err != nil {
			t.Error(err)
		}

		if req.URL.RawQuery != "" {
			t.Error("RawQuery should be an empty string")
		}

		if len(Redirects) != len(via) {
			t.Error("Redirects should be saved, lenght of Redirects should match the length via but got", len(Redirects), "and", len(via))
		}
	})

	t.Run("without RemoveQuery", func(t *testing.T) {
		RemoveQuery = false

		var req = http.Request{
			URL: &url.URL{
				RawQuery: "?id=1",
			},
		}
		
		err := logRedirect(&req, via)
		if err != nil {
			t.Error(err)
		}

		if req.URL.RawQuery != "?id=1" {
			t.Error("RawQuery should be be untouched but got", req.URL.RawQuery, "expected:", "?id=1")
		}

		if len(Redirects) != len(via) {
			t.Error("Redirects should be saved, lenght of Redirects should match the length via but got", len(Redirects), "and", len(via))
		}
	})
}