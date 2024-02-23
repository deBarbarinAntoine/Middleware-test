package router

import (
	"Middleware-test/controllers"
	"Middleware-test/internal/middlewares"
	"net/http"
)

var Mux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("GET /{$}", middlewares.Join(controllers.IndexHandlerGet, middlewares.Log()))
	Mux.HandleFunc("POST /{$}", middlewares.Join(controllers.IndexHandlerPost, middlewares.Log()))
	Mux.HandleFunc("PUT /{$}", middlewares.Join(controllers.IndexHandlerPut, middlewares.Log(), middlewares.Guard()))
	Mux.HandleFunc("DELETE /{$}", middlewares.Join(controllers.IndexHandlerDelete, middlewares.Log(), middlewares.Guard(), middlewares.Foo()))

	// Handling MethodNotAllowed error on /
	Mux.HandleFunc("/{$}", middlewares.Join(controllers.IndexHandlerNoMeth, middlewares.Log(), middlewares.Foo()))

	// Handling StatusNotFound error everywhere else
	Mux.HandleFunc("/", middlewares.Join(controllers.IndexHandlerOther, middlewares.Log(), middlewares.Foo()))
}
