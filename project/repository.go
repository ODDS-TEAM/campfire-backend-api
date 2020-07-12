package project

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	var proj Project
	if err := json.NewDecoder(r.Body).Decode(&proj); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	db, err := sql.Open("sqlite3", "../campfire.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO projects (name, description) VALUES(?, ?);", proj.Name, proj.Description)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}
}
