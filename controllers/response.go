package controllers

import (
	"encoding/json"
	"net/http"
)

// ResponseWithError is the method for error message
func ResponseWithError(w http.ResponseWriter, code int, message interface{}) {
	ResponseWithJSON(w, code, map[string]interface{}{"errors": message})
}

// ResponseWithJSON is the method for response error using json
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
