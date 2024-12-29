package handlers

import (
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
)

type viewsHandler struct{}

func NewViewsHandler() *viewsHandler {
	return &viewsHandler{}
}

func (h *viewsHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", h.homeView)
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}
