package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePong(t *testing.T) {

	handler := NewPongHandler()

	t.Run("should return 200 and pong message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/pong", nil)
		res := httptest.NewRecorder()

		handler.HandlePong(res, req)

		expectedCode := 200
		expectedBody := `{"message":"pong"}`
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
