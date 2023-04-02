package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth/cryptage"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth/security"
	"strings"
)

func ValidateToken(tok, addr string) (string, bool) {
	actualTok, _ := base64.StdEncoding.DecodeString(tok)
	res := cryptage.Decrypt(string(actualTok), security.LevelDefault)
	split := strings.Split(res, "-")

	usr := split[0]
	val := split[1]
	addr2 := split[2]

	// TEMP: we should use the stored validator in the user db.
	var validator = val
	if cryptage.Decrypt(val, security.LevelDefault) != cryptage.Decrypt(validator, security.LevelDefault) ||
		addr != cryptage.Decrypt(addr2, security.LevelDefault) {
		return "", false
	}

	return usr, true
}

func GenerateToken(username, passphrase, remoteAddr string) string {
	validator := cryptage.Encrypt(passphrase, security.LevelDefault)
	addr := cryptage.Encrypt(remoteAddr, security.LevelDefault)
	tok := cryptage.Encrypt(fmt.Sprintf("%s-%s-%s", username, validator, addr), security.LevelDefault)
	return base64.StdEncoding.EncodeToString([]byte(tok))
}
