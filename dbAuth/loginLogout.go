package dbAuth

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"authenticateAndGo/utils"

	"golang.org/x/crypto/bcrypt"
)

const hashCost = 10

var pwdDB string

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(response, nil)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"

	if !utils.IsEmpty(name) && !utils.IsEmpty(pass) {

		result := db.QueryRow("select password from users where username=$1", name)

		switch err := result.Scan(&pwdDB); err {

		case sql.ErrNoRows:
			http.Redirect(response, request, "/", 302)
		case nil:
			pwdDB = pwdDB
		default:
			http.Redirect(response, request, "/", 302)

		}

		if err := bcrypt.CompareHashAndPassword([]byte(pwdDB), []byte(pass)); err != nil {

			http.Redirect(response, request, "/", 302)

		} else {

			SetCookie(name, response)
			redirectTarget = "/index"

		}
	}

	http.Redirect(response, request, redirectTarget, 302)
}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}
