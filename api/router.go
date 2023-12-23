package api

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger" // Import the swagger package
	"net/http"
	"share-files/api/handlers"
	_ "share-files/docs" // Import your generated docs package
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.HandleDownload)
	r.Post("/upload", handlers.HandleUpload)
	r.NotFound(handlers.NotFoundHandler)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	return r
}
