package auth

import (
	"errors"
	"net/http"
	"strings"
)

func Authenticate(w http.ResponseWriter, username, passphrase, remoteAddr string) error {
	tok := GenerateToken(username, passphrase, remoteAddr)
	_, valid := ValidateToken(tok, remoteAddr)
	if !valid {
		return errors.New("couldn't validate token")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "VIMP-TKN-01",
		Value:    tok,
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

func Authenticated(r *http.Request) (string, bool) {
	c, err := r.Cookie("VIMP-TKN-01")
	if err != nil {
		return "", false
	}
	return ValidateToken(c.Value, strings.Split(r.RemoteAddr, ":")[0])
}
