package main

import (
	"github.com/shadow-wallet/bitcoind"
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/app"
	"log"
	"os"
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

	os.Setenv("BTC_D_ADDR", "localhost:8334")
	os.Setenv("BTC_D_USER", "wow")
	os.Setenv("BTC_D_PSWD", "cool")

	addr, ok := os.LookupEnv("BTC_D_ADDR")
	if !ok {
		log.Fatalln("environment variable 'BTC_D_ADDR' not set")
	}
	user, ok := os.LookupEnv("BTC_D_USER")
	if !ok {
		log.Fatalln("environment variable 'BTC_D_USER' not set")
	}
	pswd, ok := os.LookupEnv("BTC_D_PSWD")
	if !ok {
		log.Fatalln("environment variable 'BTC_D_PSWD' not set")
	}

	bc, err := bitcoind.New(addr, user, pswd)
	if err != nil {
		log.Fatalln(err)
	}
	ent, err := os.ReadDir("/home/prague/.bitcoin/wallets")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loading %d wallets\n", len(ent))
	for n, f := range ent {
		if f.IsDir() {
			err := bc.LoadWallet(f.Name())
			if err != nil {
				log.Printf("Couldn't load wallet of %s: %s", f.Name(), err)
				continue
			}
			log.Printf("Loaded wallet: %s (%d/%d)", f.Name(), n+1, len(ent))
		}
	}
	a := app.NewApp(bc)
	a.ListenAndServe(":8080")
}
