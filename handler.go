package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request is a JSON request struct
type Request interface {
	// Validate validates the request
	Validate() error
	// Handle handles the request
	Handle(w http.ResponseWriter)
}

// MakeHandler returns a handler function to serve a request
func MakeHandler(factory func() Request) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := factory()
		// Decode JSON
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteError(w, fmt.Errorf("Invalid JSON: %v", err))
			return
		}
		// Validate request
		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteError(w, fmt.Errorf("Invalid request: %v", err))
			return
		}
		// Handle panics just in case
		defer func() {
			if e := recover(); e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				WriteError(w, fmt.Errorf("Internal server error"))
			}
		}()
		req.Handle(w)
	}
}
