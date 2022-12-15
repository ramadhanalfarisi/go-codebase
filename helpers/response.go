package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message []string     `json:"message"`
}

type ResponseData struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message []string     `json:"message"`
	Data    interface{}  `json:"data"`
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

func (r *ResponseData) SendResponse(w http.ResponseWriter) {
	json, err := json.Marshal(r)
	if err != nil {
		Error(err)
	} else {
		w.Write(json)
		w.WriteHeader(r.Code)
	}
}
