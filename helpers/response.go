package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int          `json:"code,omitempty"`
	Status  string       `json:"status,omitempty"`
	Message []string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *interface{} `json:"meta,omitempty"`
}

func (r *Response) SendResponse(w http.ResponseWriter) {
	json, err := json.Marshal(r)
	if err != nil {
		Error(err)
	} else {
		w.Write(json)
		w.WriteHeader(r.Code)
	}
}
