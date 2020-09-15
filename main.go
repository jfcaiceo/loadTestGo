package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jfcaiceo/loadTestGo/controllers"
	"github.com/gorilla/mux"

	_ "github.com/joho/godotenv/autoload"
)

const (
	host = "0.0.0.0"
	port = "9000"
)

func main() {
	r := mux.NewRouter()

	var app controllers.App

	var configDB = map[string]string{
		"name":     os.Getenv("DATABASE_NAME"),
		"username": os.Getenv("DATABASE_USERNAME"),
		"password": os.Getenv("DATABASE_PASSWORD"),
		"host":     os.Getenv("DATABASE_HOST"),
		"port":     os.Getenv("DATABASE_PORT"),
		"charset":  os.Getenv("DATABASE_CHARSET"),
		"logMode":  os.Getenv("DATABASE_LOGMODE"),
	}

	app.Initialize(configDB)

	r.HandleFunc("/alive", controllers.Alive).Methods("GET")
	r.HandleFunc("/users", app.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", app.GetUser).Methods("GET")

	fmt.Println(host + ":" + port)
	http.ListenAndServe(host+":"+port, r)
}
