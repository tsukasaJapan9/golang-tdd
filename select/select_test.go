package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(delay * time.Millisecond)
				w.WriteHeader(http.StatusOK)
			}))
}

func TestRacer(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		slowServer := makeServer(20)
		defer slowServer.Close()

		fastServer := makeServer(1)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Error("expected no err")
		}

		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}

	})

	t.Run("test2", func(t *testing.T) {
		server := makeServer(100)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10 * time.Millisecond)

		if err == nil {
			t.Error("expected an error")
		}
	})
}