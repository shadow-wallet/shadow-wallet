package main

import (
	"github.com/pragu3/gophig"
	"github.com/shadow-wallet/shadow-wallet/wallet"
	"os"
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
	err := gophig.GetConfComplex("config.toml", gophig.TOMLMarshaler{}, &c)
	if os.IsNotExist(err) {
		err = gophig.SetConfComplex("config.toml", gophig.TOMLMarshaler{}, c, 777)
		if err != nil {
			panic(err)
		}
		return c
	}
	if err != nil {
		panic(err)
	}
	return c
}
