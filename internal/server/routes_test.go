package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "coffee-curator-go/internal/server"
)

func TestEntries(t *testing.T) {
	s := &server.Server{}
	t.Run("returns the journal entry with ID 1", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/entries/1", nil)
		res := httptest.NewRecorder()

		s.RegisterRoutes().ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "Coffee 1")
	})

	t.Run("create a journal entry", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/entries", nil)
		res := httptest.NewRecorder()

		s.RegisterRoutes().ServeHTTP(res, req)
		assertStatus(t, res.Code, http.StatusCreated)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response status is wrong, got %v, want %v", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}
