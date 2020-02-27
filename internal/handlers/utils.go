package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getBookId(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}
