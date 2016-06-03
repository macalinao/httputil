package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyReq struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func (r *dummyReq) Validate() error {
	if r.A == "magic" {
		return fmt.Errorf("yolo")
	}
	return nil
}

func (r *dummyReq) Handle(w http.ResponseWriter) {
	WriteJSON(w, "lol")
}

func newDummyReq() Request {
	return &dummyReq{}
}

func TestMakeHandlerHappyPath(t *testing.T) {
	data, _ := json.Marshal(&dummyReq{"test", 1})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "_", bytes.NewReader(data))

	handler := MakeHandler(newDummyReq)
	handler(w, r)

	assert.Equal(t, "application/json", w.HeaderMap.Get("content-type"))
	assert.True(t, w.Code == 200)
	assert.Equal(t, `"lol"`, w.Body.String())
}

func TestMakeHandlerInvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "_", bytes.NewReader([]byte("{brokendreams}")))

	handler := MakeHandler(newDummyReq)
	handler(w, r)

	assert.Equal(t, "application/json", w.HeaderMap.Get("content-type"))
	assert.True(t, w.Code == 400)
	assert.True(t, strings.Contains(w.Body.String(), "Invalid JSON"))
}

func TestMakeHandlerInvalidRequest(t *testing.T) {
	data, _ := json.Marshal(&dummyReq{"magic", 1})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "_", bytes.NewReader(data))

	handler := MakeHandler(newDummyReq)
	handler(w, r)

	assert.Equal(t, "application/json", w.HeaderMap.Get("content-type"))
	assert.True(t, w.Code == 400)
	assert.True(t, strings.Contains(w.Body.String(), "Invalid request"))
}
