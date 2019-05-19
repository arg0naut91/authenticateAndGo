package main

import (
	"authenticateAndGo/dbAuth"
	"authenticateAndGo/socialAuth"
	"authenticateAndGo/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {

	dbAuth.InitDB()

	router.HandleFunc("/", dbAuth.LoginPageHandler)
	router.HandleFunc("/index", dbAuth.IndexPageHandler)

	router.HandleFunc("/login", dbAuth.LoginHandler).Methods("POST")
	router.HandleFunc("/register", dbAuth.RegisterHandler).Methods("POST")

	router.HandleFunc("/loginGoogle", socialAuth.HandleGoogleLogin)
	router.HandleFunc("/loginFB", socialAuth.HandleFacebookLogin)
	router.HandleFunc("/loginGH", socialAuth.HandleGHLogin)

	router.HandleFunc("/callback", socialAuth.HandleGoogleCallback)
	router.HandleFunc("/callbackFB", socialAuth.HandleFacebookCallback)
	router.HandleFunc("/callbackGH", socialAuth.HandleGHCallback)

	router.HandleFunc("/logout", dbAuth.LogoutHandler).Methods("POST")

	fileServer := http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/")))
	router.PathPrefix("/templates/").Handler(fileServer)

	http.Handle("/", router)
	http.ListenAndServe(utils.GetPort(), router)
}
