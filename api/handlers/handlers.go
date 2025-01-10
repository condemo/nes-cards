package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	apiErrors "github.com/condemo/nes-cards/api/api_errors"
	"github.com/condemo/nes-cards/public/views/components"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func makeHandler(f CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if apiErr, ok := err.(apiErrors.ApiError); ok {
				w.WriteHeader(apiErr.Status)
				RenderTempl(w, r, components.ApiError(apiErr))
			} else {
				w.Header().Add("HX-Redirect", "/error")
				w.WriteHeader(http.StatusInternalServerError)
			}
			slog.Error("API ERROR", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
