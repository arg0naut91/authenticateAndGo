package dbAuth

import (
	"authenticateAndGo/utils"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)

	nameTempl := struct{ Name string }{userName}

	if !utils.IsEmpty(userName) {

		t, err := template.ParseFiles("./templates/success.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(response, nameTempl)

	} else {
		http.Redirect(response, request, "/", 302)
	}
}

func SetCookie(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
