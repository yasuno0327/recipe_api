package gateway

import (
	"encoding/json"
	"io"
	"net/http"
)

type Error struct {
	Message string
}

func WriteError(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(Error{Message: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)
}

func DecodeBody(domain interface{}, body io.Reader) interface{} {
	json.NewDecoder(body).Decode(domain)
	return domain
}
