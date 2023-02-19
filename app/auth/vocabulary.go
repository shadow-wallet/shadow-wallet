package auth

import (
	"github.com/pragu3/gophig"
	"log"
)

var (
	adjectives []string
	nouns      []string
	cities     []string
	verbs      []string
)

func init() {
	var voc map[string][]string
	err := gophig.GetConfComplex("assets/vocabulary.json", gophig.JSONMarshaler{}, &voc)
	if err != nil {
		log.Fatalln(err)
	}

	adjectives = voc["adjectives"]
	nouns = voc["nouns"]
	cities = voc["cities"]
	verbs = voc["verbs"]
}
