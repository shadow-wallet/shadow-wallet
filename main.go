package main

import (
	"github.com/shadow-wallet/shadow-wallet/wallet"
)

func main() {
	conf := readConfig()
	a := wallet.New(conf)
	err := a.ListenAndServe(":8080")
	if err != nil {
		panic(err)
	}
}

func readConfig() wallet.Config {
	c := wallet.DefaultConfig()
	return c
}
