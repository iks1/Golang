package main 

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
)

type StubPlayerStore struct{
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T){
    store := StubPlayerStore{
		map[string]int{
			"pepper":20,
			"floyd":10,
		},
		nil,
	}
	
	server := &PlayerServer{&store}

	t.Run("returns pepper's score", func(t *testing.T){
		request := newGetScoreRequest("pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

        assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns floyd's score", func(t *testing.T){
		request := newGetScoreRequest("floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T){
		request := newGetScoreRequest("apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		assertStatusCode(t, got, want)
	})
}

func TestStoreWins(t *testing.T){
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T){
		player := "pepper"

		request := newPostWinRequest("pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store the correct winner got %s, want %s", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
			t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertStatusCode(t testing.TB, got, want int){
	t.Helper()
	if got != want {
			t.Errorf("did not get the correct status code, got %d, want %d", got, want)
	}
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}