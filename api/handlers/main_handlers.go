package handlers

import (
	"net/http"
	"share-files/logger"
)

// HandleDownload handles GET requests and returns a welcome message.
// @Summary Returns a welcome message
// @Description This endpoint returns a welcome message.
// @ID get-welcome-message
// @Produce json
// @Success 200 {string} string "Welcome to my app!"
// @Router / [get]
func HandleDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my app!"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	logger.Log.Error().Msg("Page not found")
}
