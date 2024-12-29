package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type ApiError struct {
	Err    error
	Msg    string
	Status int
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

func NewApiError(err error, msg string, status int) ApiError {
	return ApiError{Err: err, Msg: msg, Status: status}
}

var internalError = map[string]any{
	"status": http.StatusInternalServerError,
	"msg":    "internal server error",
}

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func makeHandler(f CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if apiErr, ok := err.(ApiError); ok {
				fmt.Fprintf(w, "unformat err - %v", apiErr)
			} else {
				fmt.Fprint(w, internalError)
			}
			slog.Error("API ERROR", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
