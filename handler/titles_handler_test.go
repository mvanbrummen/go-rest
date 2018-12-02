package handler

import (
	"net/http"
	"testing"
)

func TestGetTitle(t *testing.T) {
	req, err := http.NewRequest("GET", "/titles/poo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// handler := NewTitlesHandler()

	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(HealthCheckHandler)

	// handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
