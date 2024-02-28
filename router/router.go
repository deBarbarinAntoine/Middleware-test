package router

import (
	"Middleware-test/controllers"
	"net/http"
)

var Mux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("GET /{$}", controllers.IndexHandlerGetBundle)
	Mux.HandleFunc("PUT /{$}", controllers.IndexHandlerPutBundle)
	Mux.HandleFunc("DELETE /{$}", controllers.IndexHandlerDeleteBundle)
	Mux.HandleFunc("POST /add", controllers.IndexHandlerPutBundle)
	Mux.HandleFunc("GET /login", controllers.LoginHandlerGetBundle)
	Mux.HandleFunc("POST /login", controllers.LoginHandlerPostBundle)
	Mux.HandleFunc("GET /register", controllers.RegisterHandlerGetBundle)
	Mux.HandleFunc("POST /register", controllers.RegisterHandlerPostBundle)
	Mux.HandleFunc("GET /home", controllers.HomeHandlerGetBundle)
	Mux.HandleFunc("GET /logs", controllers.LogHandlerGetBundle)

	// Handling MethodNotAllowed error on /
	Mux.HandleFunc("/{$}", controllers.IndexHandlerNoMethBundle)

	// Handling StatusNotFound error everywhere else
	Mux.HandleFunc("/", controllers.IndexHandlerOtherBundle)
}
