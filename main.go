package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}

	srv.ListenAndServe()
}
