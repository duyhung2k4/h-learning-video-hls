package router

import (
	"app/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func apiV1(router chi.Router) {

	fileController := controller.NewFileController()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"mess": "done",
		})
	})

	router.Route("/public", func(public chi.Router) {
	})

	router.Route("/video", func(encoding chi.Router) {
		encoding.Get("/{uuid}/{quantity}/{filename}", fileController.FileHls)
	})

}
