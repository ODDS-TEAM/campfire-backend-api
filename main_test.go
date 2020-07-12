package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func clearTables() {
	db, _ := sql.Open("sqlite3", "campfire.db")
	db.Exec("DELETE FROM projects")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='projects';")
}

func TestCreateNewProject(t *testing.T) {
	clearTables()

	expectedName := "My Wealth"
	expectedDescription := "TISCO Mobile Banking"

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&Project{Name: expectedName, Description: expectedDescription})
	req, _ := http.NewRequest("POST", "/projects", payload)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(ProjectHandler)
	handler.ServeHTTP(w, req)

	status := w.Code
	assert.Equal(t, status, http.StatusOK, "Handler returned wrong status code")

	db, _ := sql.Open("sqlite3", "campfire.db")
	sqlStatement := `SELECT name, description FROM projects WHERE id=$1;`
	row := db.QueryRow(sqlStatement, 1)

	var name string
	var description string
	if err := row.Scan(&name, &description); err == nil {
		assert.Equal(t, expectedName, name)
		assert.Equal(t, expectedDescription, description)
	}
}
