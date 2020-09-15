package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfcaiceo/loadTestGo/config"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB      *gorm.DB
}

func (app *App) Initialize(dbConfig map[string]string) {
	app.DB = config.DB(dbConfig)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	fmt.Print(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
