package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	server "coffee-curator-go/internal/server"
)

func TestHelloWorldHandler(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	assertStatus(t, resp.StatusCode, http.StatusOK)
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	assertResponseBody(t, expected, string(body))
}

func TestEntries(t *testing.T) {
	t.Run("returns the journal entry with ID 1", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/entries/1", nil)
		res := httptest.NewRecorder()

		s := &server.Server{}
		s.RegisterRoutes().ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "Coffee 1" )
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
