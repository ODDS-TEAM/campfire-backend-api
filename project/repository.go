package project

import (
	"fmt"
	"net/http"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
