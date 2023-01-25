package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:3001/auth/google/callback", // Oauth가 처리가 끝난뒤 결과를 callback 해줄곳
	ClientID:     "791518846100-sdlfvc03tcmj0987eaq8ni1ltpg7pldb.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, // https://developers.google.com/identity/protocols/oauth2/scopes?hl=ko > scope 문서
	Endpoint:     google.Endpoint,
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	// cookie에 temp key를 심고, 쿠키를 비교해본다.
	state := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b) // byte 배열을 랜덤하게 채움

	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)

	return state
}

func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprint(w, string(data))
}

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)

	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s", err.Error())
	}

	return ioutil.ReadAll(resp.Body)
}
