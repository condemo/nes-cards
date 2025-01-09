package middlewares

import (
	"fmt"
	"net/http"
)

type Middleware func(next http.Handler) http.HandlerFunc

func MiddlewareStack(m ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next.ServeHTTP
	}
}

func Recover(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: Implementar esto con una vista o notificación; más acorde con HTMX
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	}
}
