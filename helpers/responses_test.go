package helpers

import (
	"assessment/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccess(t *testing.T) {
	InitializeLogger(logger.NewTestLogger())
	var codes = []int{200, 201, 202, 204}
	handler := func(w http.ResponseWriter, r *http.Request, code int) {
		ReturnSuccess(r, w, code, "Test Returned Success Successfully")
	}

	req := httptest.NewRequest("GET", "https://area99test.net", nil)

	for _, code := range codes {
		w := httptest.NewRecorder()
		handler(w, req, code)

		resp := w.Result()

		if resp.StatusCode != code {
			t.Errorf("Expected: %v, Got: %v\n", code, resp.StatusCode)
		}
	}
}

func TestFailure(t *testing.T) {
	InitializeLogger(logger.NewTestLogger())
	var codes = []int{400, 401, 402, 403, 404, 405, 500, 503, 504}
	handler := func(w http.ResponseWriter, r *http.Request, code int) {
		ReturnFailure(r, w, code, "An Error Occurred")
	}

	req := httptest.NewRequest("GET", "https://area99test.net", nil)

	for _, code := range codes {
		w := httptest.NewRecorder()
		handler(w, req, code)

		resp := w.Result()

		if resp.StatusCode != code {
			t.Errorf("Expected: %v, Got: %v\n", code, resp.StatusCode)
		}
	}
}
