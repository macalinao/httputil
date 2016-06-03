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
	Error string `json:"error"`
}

// WriteError writes an error to an http.ResponseWriter
func WriteError(w http.ResponseWriter, msg interface{}) {
	WriteJSON(w, httpError{fmt.Sprintf("%v", msg)})
}
