package app

import (
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/app/auth"
	"net/http"
	"strings"
)

func (a *App) Generate(w http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Authenticated(r); ok {
		redirectHome(w, r)
		return
	}
	username := auth.GenerateUsername()
	err := a.bc.CreateWallet(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = a.bc.GetNewAddress(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	passphrase := auth.GeneratePassphrase()

	err = a.bc.EncryptWallet(username, passphrase)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = auth.Authenticate(w, username, passphrase, strings.Split(r.RemoteAddr, ":")[0])
	if err != nil {
		_, _ = fmt.Fprint(w, err)
	}
	_, _ = fmt.Fprintf(w, "Generated username: %s\nGenerated passphrase: %s\nsave these information", username, passphrase)
}
