package auth

import (
	"encoding/base64"
	"fmt"
	ars2 "github.com/shadow-wallet/shadow-wallet/wallet/auth/crypto/ars"
	"strings"
)

func ValidateToken(tok, addr string) (string, bool) {
	actualTok, _ := base64.StdEncoding.DecodeString(tok)
	res := ars2.Decrypt(string(actualTok), ars2.LevelDefault)
	split := strings.Split(res, "-")

	usr := split[0]
	val := split[1]
	addr2 := split[2]

	// TEMP: we should use the stored validator in the user db.
	var validator = val
	if ars2.Decrypt(val, ars2.LevelDefault) != ars2.Decrypt(validator, ars2.LevelDefault) ||
		addr != ars2.Decrypt(addr2, ars2.LevelDefault) {
		return "", false
	}

	return usr, true
}

func GenerateToken(username, passphrase, remoteAddr string) string {
	validator := ars2.Encrypt(passphrase, ars2.LevelDefault)
	addr := ars2.Encrypt(remoteAddr, ars2.LevelDefault)
	tok := ars2.Encrypt(fmt.Sprintf("%s-%s-%s", username, validator, addr), ars2.LevelDefault)
	return base64.StdEncoding.EncodeToString([]byte(tok))
}
