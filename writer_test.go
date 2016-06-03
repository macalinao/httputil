package httputil

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	data := map[string]interface{}{
		"a": "hi",
		"b": 1,
	}
	w := httptest.NewRecorder()
	WriteJSON(w, data)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, `{"a":"hi","b":1}`, w.Body.String())
}

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	WriteError(w, "you dun goofed")

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, `{"success":false,"error":"you dun goofed"}`, w.Body.String())
}
