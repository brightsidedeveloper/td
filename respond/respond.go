package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

type Responder interface {
	Error(w http.ResponseWriter, status int, message string)
	JSON(w http.ResponseWriter, status int, payload interface{})
}

type responder struct {}

func NewResponder() responder {
	return responder{}
}


func (r responder) Error(w http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Printf("Responding with 5XX error: %v", message)

	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	r.JSON(w, status, ErrorResponse{Error: message})
}

func (r responder) JSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}