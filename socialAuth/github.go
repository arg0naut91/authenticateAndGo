package socialAuth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GithubUser struct {
	Login  string `json:"login"`
	NodeID string `json:"node_id"`
	GHName string `json:"name"`
	Email  string `json:"email"`
}

var (
	ghOauthConf = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callbackGH",
		ClientID:     os.Getenv("GH_CLIENT_ID"),
		ClientSecret: os.Getenv("GH_CLIENT_SECRET"),
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
)

func HandleGHLogin(w http.ResponseWriter, r *http.Request) {

	url := ghOauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func HandleGHCallback(w http.ResponseWriter, r *http.Request) {

	user, err := GetUserInfoGH(r.FormValue("state"), r.FormValue("code"))

	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	us := string(user.GHName)

	GithubName := struct{ Name string }{strings.Fields(us)[0]}

	t, err := template.ParseFiles("./templates/success.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, GithubName)
}

func GetUserInfoGH(state string, code string) (*GithubUser, error) {

	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := GetAccessToken(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := GetOauthUser(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return response, nil
}

func GetAccessToken(ctx context.Context, code string) (string, error) {
	client := HTTPClient()
	data, err := json.Marshal(map[string]interface{}{
		"client_id":     os.Getenv("GH_CLIENT_ID"),
		"client_secret": os.Getenv("GH_CLIENT_SECRET"),
		"code":          code,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}
	if body["error"] != "" {
		return "", fmt.Errorf("Error in response body: %s", err.Error())
	}
	return body["access_token"], nil
}

func GetOauthUser(ctx context.Context, accessToken string) (*GithubUser, error) {
	client := HTTPClient()
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user GithubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	email, err := GetEmail(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	user.Email = email
	return &user, nil
}

func GetEmail(ctx context.Context, accessToken string) (string, error) {
	client := HTTPClient()
	req, err := http.NewRequest("GET", "https://api.github.com/user/public_emails", nil)
	if err != nil {
		return "", err
	}
	req.Close = true
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emails []struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return "", err
	}
	if len(emails) == 0 {
		return "", nil
	}
	return emails[0].Email, nil
}

func HTTPClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}
