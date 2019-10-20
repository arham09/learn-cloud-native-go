package router

import (
	"github.com/go-chi/chi"

	"go-app/app/app"
	"go-app/app/handler"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", handler.NewHandler(a.HandleIndex, l))

	return r
}
