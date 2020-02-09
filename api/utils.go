package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
	"github.com/go-chi/render"
)

type ErrInvalidParameter struct {
	Code   string
	Name   string
	Reason string
}

func CreateErrInvalidParameter(name string, reason string) ErrInvalidParameter {
	return ErrInvalidParameter{Code: "INVALID_PARAMETER", Name: name, Reason: reason}
}

func (e *ErrInvalidParameter) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 400)
	return nil
}

// Conjson is a JSON serializer that uses reflection to automagically turn
// FooBar -> fooBar.
func RenderConjson(w http.ResponseWriter, r *http.Request, v interface{}) {
	buf := &bytes.Buffer{}
	jsonEnc := json.NewEncoder(buf)
	jsonEnc.SetEscapeHTML(true)
	enc := conjson.NewEncoder(jsonEnc, transform.CamelCaseKeys(true))
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write(buf.Bytes())
}
