package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r* http.Request) {
	data := map[string]string {
		"status": "ok", 
		"env": app.config.env,
		"version": version,
	}
	if err := WriteJson(w, http.StatusOK, data); err != nil {
		log.Println(err.Error())
	}
}