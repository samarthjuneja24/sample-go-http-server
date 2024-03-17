package routes

import (
	"net/http"
)

// SetupRouter sets up the router for the HTTP server.
func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello world!"))
	})

	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Status-ok"))
	})
	return router
}
