package auth

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateUsername() string {
	rand.Seed(time.Now().UnixNano())
	a, n := rand.Intn(len(adjectives)-1), rand.Intn(len(nouns)-1)
	num := rand.Intn(9999)

	return fmt.Sprintf("%s%s%d", adjectives[a], nouns[n], num)
}

func GeneratePassphrase() string {
	rand.Seed(time.Now().UnixNano())
	var phrase []string

	for len(phrase) < 24 {
		if n := rand.Intn(5); n == 1 {
			phrase = append(phrase, adjectives[rand.Intn(len(adjectives)-1)])
		} else if n == 2 {
			phrase = append(phrase, nouns[rand.Intn(len(nouns)-1)])
		} else if n == 3 {
			phrase = append(phrase, cities[rand.Intn(len(cities)-1)])
		} else {
			phrase = append(phrase, verbs[rand.Intn(len(verbs)-1)])
		}
	}
	return strings.Join(phrase, " ")
}
