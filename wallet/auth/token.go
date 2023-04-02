package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth/ars"
	"strings"
)

func ValidateToken(tok, addr string) (string, bool) {
	actualTok, _ := base64.StdEncoding.DecodeString(tok)
	res := ars.Decrypt(string(actualTok), ars.LevelDefault)
	split := strings.Split(res, "-")

	usr := split[0]
	val := split[1]
	addr2 := split[2]

	// TEMP: we should use the stored validator in the user db.
	var validator = val
	if ars.Decrypt(val, ars.LevelDefault) != ars.Decrypt(validator, ars.LevelDefault) ||
		addr != ars.Decrypt(addr2, ars.LevelDefault) {
		return "", false
	}

	return usr, true
}

func GenerateToken(username, passphrase, remoteAddr string) string {
	validator := ars.Encrypt(passphrase, ars.LevelDefault)
	addr := ars.Encrypt(remoteAddr, ars.LevelDefault)
	tok := ars.Encrypt(fmt.Sprintf("%s-%s-%s", username, validator, addr), ars.LevelDefault)
	return base64.StdEncoding.EncodeToString([]byte(tok))
}
