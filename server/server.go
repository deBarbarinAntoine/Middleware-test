package server

import (
	"Middleware-test/router"
	"log"
	"net/http"
)

func Run() {
	router.Init()
	fs := http.FileServer(http.Dir("../assets"))
	router.Mux.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatalln(http.ListenAndServe(":8080", router.Mux))
}
