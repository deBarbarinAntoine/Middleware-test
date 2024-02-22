package controllers

import (
	"Middleware-test/internal/middlewares"
	"Middleware-test/internal/utils"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

func IndexHandlerGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerGet")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandlerPost(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerPost")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandlerPut(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerPut")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandlerDelete(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerDelete")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandlerNoMeth(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Error", http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	middlewares.Logger.Warn("IndexHandlerNoMeth", slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusMethodNotAllowed))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusMethodNotAllowed) + " !"))
}

func IndexHandlerOther(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Error", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	middlewares.Logger.Warn("IndexHandlerOther", slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusNotFound))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusNotFound) + " !"))
}
