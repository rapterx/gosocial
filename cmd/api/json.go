package main

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	//only allow body of less than 1mb
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decocder := json.NewDecoder(r.Body)
	decocder.DisallowUnknownFields()

	return decocder.Decode(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return WriteJson(w, status, &envelope{Error: message})
}