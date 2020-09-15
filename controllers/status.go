package controllers

import "net/http"

func Alive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
