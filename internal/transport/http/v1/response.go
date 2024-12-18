package v1

import (
	"net/http"

	"github.com/go-chi/render"
)

type HttpError struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Message string        `json:"message"`
	DebugID string        `json:"debug_id,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
	Code    int           `json:"code"`
}

type ErrorDetail struct {
	Fields  []string `json:"fields,omitempty"`
	Message string   `json:"message,omitempty"`
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, obj any) {
	if obj == nil {
		obj = struct {
		}{}
	}
	render.JSON(w, r, obj)
}

func ErrorJSON(w http.ResponseWriter, r *http.Request, code int, err error, errs ...ErrorDetail) {
	var resp HttpError

	resp.Meta.Code = code
	resp.Meta.Message = err.Error()
	resp.Meta.Details = errs

	render.Status(r, code)
	render.JSON(w, r, &resp)
}
