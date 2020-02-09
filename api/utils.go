package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
	"github.com/go-chi/render"
)

type ErrResponse struct {
	StatusCode int `json:"-"`

	Error Err
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

type Err struct {
	Code        string
	Description string
	Detail      interface{}
}

type InvalidParameterDetail struct {
	Name   string
	Reason string
}

func CreateErrInvalidParameter(name string, reason string) ErrResponse {
	return ErrResponse{
		StatusCode: 400,
		Error: Err{
			Code:        "INVALID_PARAMETER",
			Description: "Could not parse invalid parameter",
			Detail: InvalidParameterDetail{
				Name:   name,
				Reason: reason,
			},
		},
	}
}

func CreateErrNotFound(resourceName string) ErrResponse {
	return ErrResponse{
		StatusCode: 404,
		Error: Err{
			Code:        "NOT_FOUND",
			Description: fmt.Sprintf("%s not found", resourceName),
		},
	}
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
