package ars

import (
	"encoding/base64"
	"strings"
)

// Encrypt encrypts the given string according to the given security level.
func Encrypt(str string, security Level) string {
	newStr := ""
	str = base64.StdEncoding.EncodeToString([]byte(str))
	for _, i := range str {
		newStr += string(encryptChar(i, security.Spacing))
	}
	return newStr + security.FirstStrBreak + newStr[:1] + security.SecondStrBreak
}

// Decrypt decrypts the given string according to the given security level.
func Decrypt(str string, security Level) string {
	newStr := ""
	hash := strings.Split(str, "`")[0]
	for _, char := range hash {
		newStr += string(decryptChar(char, security.Spacing))
	}
	s, _ := base64.StdEncoding.DecodeString(newStr)
	return string(s)
}

// encryptChar encrypts the given character according to the given spacing.
func encryptChar(char rune, spacing int) rune {
	var newChar rune
	for i, ch := range characters {
		if ch == char {
			if i >= spacing {
				for true {
					if i+spacing <= len(characters)-1 {
						newChar = characters[i+spacing]
						break
					} else if (i+spacing)-len(characters) <= len(characters)-1 {
						newChar = characters[(i+spacing)-len(characters)]
						break
					}
					break
				}
			} else if i < spacing {
				for true {
					if i+spacing <= len(characters)-1 {
						newChar = characters[i+spacing]
						break
					}
					break
				}
			}
		}
	}
	return newChar
}

// decryptChar decrypts the given character according to the given spacing.
func decryptChar(char rune, spacing int) rune {
	var newChar rune
	for i, ch := range characters {
		if ch == char {
			if i >= spacing {
				for true {
					if i-spacing <= len(characters)-1 {
						newChar = characters[i-spacing]
						break
					}
					break
				}
			} else if i < spacing {
				for true {
					if i-spacing <= len(characters)-1 {
						newChar = characters[i-spacing]
						break
					} else if (i-spacing)+len(characters) <= len(characters)-1 {
						newChar = characters[(i-spacing)+len(characters)]
						break
					}
					break
				}
			}
		}
	}
	return newChar
}
