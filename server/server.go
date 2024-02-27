package server

import (
	"Middleware-test/internal/utils"
	"Middleware-test/router"
	"log"
	"net/http"
)

func Run() {
	router.Init()
	fs := http.FileServer(http.Dir("../assets"))
	router.Mux.Handle("/static/", http.StripPrefix("/static/", fs))
	go utils.MonitorSessions()
	log.Fatalln(http.ListenAndServe(":8080", router.Mux))
}
