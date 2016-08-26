package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSON writes JSON to an http.ResponseWriter
func WriteJSON(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

type httpError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// WriteError writes an error to an http.ResponseWriter
func WriteError(w http.ResponseWriter, msg interface{}) {
	WriteJSON(w, httpError{false, fmt.Sprintf("%v", msg)})
}

// WriteNotFound writes not found
func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	WriteError(w, "Not found")
}
