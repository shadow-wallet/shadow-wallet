package app

import (
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/app/auth"
	"log"
	"net/http"
	"strings"
)

type loginInfo struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Authenticated(r); ok {
		redirectHome(w, r)
		return
	}

	if r.Method == http.MethodGet {
		loginTemplate.Execute(w, nil)
		return
	}

	var lgInfo loginInfo
	_, err := unmarshalQuery(r.Body, &lgInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = auth.Authenticate(w, lgInfo.Username, lgInfo.Passphrase, strings.Split(r.RemoteAddr, ":")[0])
	if err != nil {
		_, _ = fmt.Fprint(w, err)
	}
	redirectHome(w, r)
}
