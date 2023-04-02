package wallet

import (
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth"
	"net/http"
	"strings"
)

func (w *Wallet) Generate(wr http.ResponseWriter, r *http.Request) {
	if _, ok := auth.Authenticated(r); ok {
		w.redirectHome(wr, r)
		return
	}
	username := auth.GenerateUsername()
	err := w.bc.CreateWallet(username)
	if err != nil {
		wr.Write([]byte(err.Error()))
		return
	}
	_, err = w.bc.GetNewAddress(username)
	if err != nil {
		wr.Write([]byte(err.Error()))
		return
	}
	passphrase := auth.GeneratePassphrase()

	err = w.bc.EncryptWallet(username, passphrase)
	if err != nil {
		wr.Write([]byte(err.Error()))
		return
	}

	err = auth.Authenticate(wr, username, passphrase, strings.Split(r.RemoteAddr, ":")[0])
	if err != nil {
		_, _ = fmt.Fprint(wr, err)
	}
	_, _ = fmt.Fprintf(wr, "Generated username: %s\nGenerated passphrase: %s\nsave these information", username, passphrase)
}
