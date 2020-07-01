package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealhCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	assert.Equal(t, status, http.StatusOK, "Handler returned wrong status code")

	expected := "ok"
	assert.Equal(t, expected, rr.Body.String(), "Handler returned unexpected body")
}
