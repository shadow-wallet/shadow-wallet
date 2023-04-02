package wallet

import (
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth"
	"log"
	"net/http"
	"strings"
)

type loginInfo struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
}

func (w *Wallet) Login(wr http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Authenticated(r); ok {
		w.redirectHome(wr, r)
		return
	}

	if r.Method == http.MethodGet {
		loginTemplate.Execute(wr, nil)
		return
	}

	var lgInfo loginInfo
	_, err := unmarshalQuery(r.Body, &lgInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = auth.Authenticate(wr, lgInfo.Username, lgInfo.Passphrase, strings.Split(r.RemoteAddr, ":")[0])
	if err != nil {
		_, _ = fmt.Fprint(wr, err)
	}
	w.redirectHome(wr, r)
}
