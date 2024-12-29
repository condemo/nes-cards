package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
