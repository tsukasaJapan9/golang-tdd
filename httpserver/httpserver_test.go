package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/players/%s", name),
		nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:%s, want:%s", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestHttpServer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int {
			"pepper":20,
			"floyd":10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

