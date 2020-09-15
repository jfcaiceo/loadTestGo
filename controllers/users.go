package controllers

import (
	"net/http"
	"strconv"

	"github.com/jfcaiceo/loadTestGo/models"
	"github.com/gorilla/mux"
)

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	
	var user models.User
	users, _ := user.All(app.DB, map[string]interface{}{"email": email})

	respondWithJSON(w, http.StatusOK, users)
}

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user := models.User{ID: uint(id)}
	if err := user.Get(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

