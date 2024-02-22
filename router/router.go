package router

import (
	"Middleware-test/controllers"
	"Middleware-test/middlewares"
	"net/http"
)

var Mux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("GET /{$}", middlewares.Log()(controllers.IndexHandlerGet))
	Mux.HandleFunc("POST /{$}", middlewares.Log()(controllers.IndexHandlerPost))
	Mux.HandleFunc("PUT /{$}", middlewares.Log()(middlewares.Guard()(controllers.IndexHandlerPut)))
	Mux.HandleFunc("DELETE /{$}", middlewares.Log()(middlewares.Guard()(controllers.IndexHandlerDelete)))

	// Handling MethodNotAllowed error on /
	Mux.HandleFunc("/{$}", middlewares.Log()(controllers.IndexHandlerNoMeth))

	// Handling StatusNotFound error everywhere else
	Mux.HandleFunc("/", middlewares.Log()(controllers.IndexHandlerOther))
}
