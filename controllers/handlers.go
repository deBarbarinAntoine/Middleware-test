package controllers

import (
	"Middleware-test/internal/middlewares"
	"Middleware-test/internal/models"
	"Middleware-test/internal/utils"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strings"
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

func indexHandlerPut(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	sessionID, _ := r.Cookie("updatedCookie")
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
	sessionID, _ := r.Cookie("updatedCookie")
	err = tmpl.ExecuteTemplate(w, "index", "indexHandlerDelete"+sessionID.Value+"\nUsername: "+utils.SessionsData[sessionID.Value].Username+"\nIP address: "+utils.SessionsData[sessionID.Value].IpAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandlerNoMeth(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	log.Println("HTTP Error", http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	utils.Logger.Warn("indexHandlerNoMeth", slog.Int("reqID", middlewares.LogId), slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusMethodNotAllowed))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusMethodNotAllowed) + " !"))
}

func indexHandlerOther(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	log.Println("HTTP Error", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	utils.Logger.Warn("indexHandlerOther", slog.Int("reqID", middlewares.LogId), slog.String("reqURL", r.URL.String()), slog.Int("HttpStatus", http.StatusNotFound))
	w.Write([]byte("Error " + fmt.Sprint(http.StatusNotFound) + " !"))
}

func loginHandlerGet(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	var message template.HTML
	if r.URL.Query().Has("err") {
		switch r.URL.Query().Get("err") {
		case "login":
			message = "<div class=\"message\">Wrong username or password!</div>"
		case "restricted":
			message = "<div class=\"message\">You need to login to access that area!</div>"
		}
	}
	tmpl, err := template.ParseFiles(utils.Path + "templates/login.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "login", message)
	if err != nil {
		log.Fatalln(err)
	}
}

func loginHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	credentials := models.Credentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	if utils.CheckPwd(credentials) {
		utils.OpenSession(&w, credentials.Username, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login?err=login", http.StatusSeeOther)
	}
}

func registerHandlerGet(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	var message template.HTML
	if r.URL.Query().Has("err") {
		switch r.URL.Query().Get("err") {
		case "username":
			message = "<div class=\"message\">Username already used!</div>"
		case "password":
			message = "<div class=\"message\">Both passwords need to be equal!</div>"
		case "email":
			message = "<div class=\"message\">Wrong email value!</div>"
		}
	}
	tmpl, err := template.ParseFiles(utils.Path + "templates/register.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "register", message)
	if err != nil {
		log.Fatalln(err)
	}
}

func registerHandlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	formValues := struct {
		username  string
		email     string
		password1 string
		password2 string
	}{
		username:  r.FormValue("username"),
		email:     strings.TrimSpace(strings.ToLower(r.FormValue("email"))),
		password1: r.FormValue("password1"),
		password2: r.FormValue("password2"),
	}
	_, exists := utils.SelectUser(formValues.username)
	switch {
	case exists:
		http.Redirect(w, r, "register?err=username", http.StatusSeeOther)
		return
	case formValues.password1 != formValues.password2:
		http.Redirect(w, r, "register?err=password", http.StatusSeeOther)
		return
	case !utils.CheckEmail(formValues.email):
		http.Redirect(w, r, "register?err=email", http.StatusSeeOther)
		return
	}
	hash, salt := utils.NewPwd(formValues.password1)
	newUser := models.User{
		Id:        utils.GetIdNewUser(),
		Username:  formValues.username,
		HashedPwd: hash,
		Salt:      salt,
		Email:     formValues.email,
	}
	utils.CreateUser(newUser)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func homeHandlerGet(w http.ResponseWriter, r *http.Request) {
	log.Println(utils.GetCurrentFuncName())
	tmpl, err := template.ParseFiles(utils.Path + "templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", "homeHandlerGet --- Restricted area! ---")
	if err != nil {
		log.Fatalln(err)
	}
}
