package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
)


type errorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func ErrorResponse(msg string, err error) errorResponse {
	return errorResponse{
		Message: msg,
		Details: err.Error(),
	}
}

func Response(w http.ResponseWriter, code int, er interface{}) {
	w.WriteHeader(code)
	JSON(w, er)
}

func JSON(w http.ResponseWriter, msg interface{}) {
	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	e.SetEscapeHTML(true)
	if err := e.Encode(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}