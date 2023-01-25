package oauth

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func OauthServerExec() {
	mux := pat.New()
	mux.HandleFunc("/auth/google/login", googleLoginHandler)
	mux.HandleFunc("/auth/google/callback", googleAuthCallback)

	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":3001", n)
}
