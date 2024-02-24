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

func indexHandlerGet(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerGet")
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	utils.OpenSession(&w, r)
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerPost")
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandlerPut(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	sessionID, _ := r.Cookie("session_id")
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerPut"+sessionID.Value+"\nUsername: "+utils.SessionsData[sessionID.Value].Username+"\nIP address: "+utils.SessionsData[sessionID.Value].IpAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandlerDelete(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	sessionID, _ := r.Cookie("session_id")
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerDelete"+sessionID.Value+"\nUsername: "+utils.SessionsData[sessionID.Value].Username+"\nIP address: "+utils.SessionsData[sessionID.Value].IpAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandlerNoMeth(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	log.Println("HTTP Error", http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	middlewares.Logger.Warn("indexHandlerNoMeth", slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusMethodNotAllowed))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusMethodNotAllowed) + " !"))
}

func indexHandlerOther(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	log.Println("HTTP Error", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	middlewares.Logger.Warn("indexHandlerOther", slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusNotFound))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusNotFound) + " !"))
}
