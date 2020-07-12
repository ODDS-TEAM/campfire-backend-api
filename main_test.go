package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealhCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(w, req)

	status := w.Code
	assert.Equal(t, status, http.StatusOK, "Handler returned wrong status code")

	expected := "ok"
	assert.Equal(t, expected, w.Body.String(), "Handler returned unexpected body")
}
