package responder

import (
	"encoding/json"
	"net/http"
)

type Responder struct {
}

func NewResponder() *Responder {
	return &Responder{}
}

func (w *Responder) JsonOk(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	if data == nil {
		data = make(map[string]interface{}, 0)
	}

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		// TODO:Убрать панику
		panic(err)
	}
}

func (w *Responder) JsonErr(writer http.ResponseWriter, status int, err error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := make(map[string]string)
	if err != nil {
		data["error"] = err.Error()
	}

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		// TODO:Убрать панику
		panic(err)
	}
}
