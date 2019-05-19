package dbAuth

import (
	"net/http"

	"authenticateAndGo/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	uName := r.FormValue("username")
	email := r.FormValue("emailsignup")
	pwd := r.FormValue("passwordsignup")
	confirmPwd := r.FormValue("passwordconfirm")

	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !utils.IsEmpty(uName)
	_email = !utils.IsEmpty(email)
	_pwd = !utils.IsEmpty(pwd)
	_confirmPwd = !utils.IsEmpty(confirmPwd)

	if _uName && _email && _pwd && _confirmPwd {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 8)

		if _, err = db.Query("insert into users values ($1, $2, $3)", uName, email, string(hashedPassword)); err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		SetCookie(uName, w)
		http.Redirect(w, r, "/index", 302)

	}
}
