package main

import (
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/wallet"
	"time"
)

func main() {
	cm := coinmarket.New("a1528619-40d3-4f53-9a24-2473d82a1660")
	cm.UpdateFiat()
	go func() {
		for {
			time.Sleep(time.Minute * 5)
			cm.UpdateFiat()
		}
	}()

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
