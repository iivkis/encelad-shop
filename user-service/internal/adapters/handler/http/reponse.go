package httphandler

import (
	"encoding/json"
	"net/http"
)

type ResponseEncoder struct {
}

func NewResponseEncoder() *ResponseEncoder {
	return &ResponseEncoder{}
}

func (w *ResponseEncoder) JsonOk(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	if data == nil {
		data = make(map[string]interface{}, 0)
	}

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}

func (w *ResponseEncoder) JsonErr(writer http.ResponseWriter, status int, err error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := make(map[string]string, 1)
	if err != nil {
		data["error"] = err.Error()
	}

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
