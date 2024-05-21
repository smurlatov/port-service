package handler

import (
	"log"
	"net/http"
)

func RespondOK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, "Bad request", http.StatusBadRequest)
}

func httpRespondWithError(err error, w http.ResponseWriter, r *http.Request, msg string, status int) {
	log.Printf("error: %s, slug: %s, msg: %s", err, msg)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
}
