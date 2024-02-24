package router

import (
	"Middleware-test/controllers"
	"net/http"
)

var Mux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("GET /{$}", controllers.IndexHandlerGetBundle)
	Mux.HandleFunc("POST /{$}", controllers.IndexHandlerPostBundle)
	Mux.HandleFunc("PUT /{$}", controllers.IndexHandlerPutBundle)
	Mux.HandleFunc("DELETE /{$}", controllers.IndexHandlerDeleteBundle)
	Mux.HandleFunc("POST /add", controllers.IndexHandlerPutBundle)

	// Handling MethodNotAllowed error on /
	Mux.HandleFunc("/{$}", controllers.IndexHandlerNoMethBundle)

	// Handling StatusNotFound error everywhere else
	Mux.HandleFunc("/", controllers.IndexHandlerOtherBundle)
}
