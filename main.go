package main

import (
	"github.com/shadow-wallet/shadow-wallet/wallet"
)

func main() {
	conf := readConfig()
	w, err := wallet.New(conf)
	if err != nil {
		panic(err)
	}
	err = w.ListenAndServe(":8080")
	if err != nil {
		panic(err)
	}
}

func readConfig() wallet.Config {
	c := wallet.DefaultConfig()
	return c
}
