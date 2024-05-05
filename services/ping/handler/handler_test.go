package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePing(t *testing.T) {

	handler := NewPingHandler()

	t.Run("should return 200 and ping message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		res := httptest.NewRecorder()

		handler.HandlePing(res, req)

		expectedCode := 200
		expectedBody := `{"message":"ping"}`
		assertStatusCode(t, res.Code, expectedCode)
		assertBody(t, res.Body.String(), expectedBody)
	})
}

func assertStatusCode(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("got status %d but expected status %d", actual, expected)
	}
}

func assertBody(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("got body %s but expected body %s", actual, expected)
	}
}
