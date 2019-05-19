package socialAuth

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthStateString = "f43jf49999f4kpoe"
)

type FBUser struct {
	FBName string `json:"first_name"`
	Email  string `json:"email"`
}

var (
	fbOauthConf = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callbackFB",
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
)

func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {

	url := fbOauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {

	content, err := GetUserInfoFacebook(r.FormValue("state"), r.FormValue("code"))

	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var user *FBUser
	_ = json.Unmarshal(content, &user)

	FacebookName := struct{ Name string }{user.FBName}

	t, err := template.ParseFiles("./templates/success.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, FacebookName)
}

func GetUserInfoFacebook(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := fbOauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://graph.facebook.com/me?fields=first_name,email&access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
