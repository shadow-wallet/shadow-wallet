package wallet

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shadow-wallet/bitcoind"
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/wallet/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	loginTemplate, _ = template.New("assets/frontend/login.html", "assets/frontend/login.css")
)

type Wallet struct {
	bc   *bitcoind.Bitcoind
	cm   *coinmarket.CoinMarket
	conf Config
}

func New(conf Config) (*Wallet, error) {
	bc, err := bitcoind.New(conf.Daemon.Address,
		conf.Daemon.Username, conf.Daemon.Password)
	if err != nil {
		return nil, err
	}
	w := &Wallet{
		cm:   coinmarket.New(conf.CoinMarket.APIKEY),
		bc:   bc,
		conf: conf,
	}
	return w, nil
}

func (w *Wallet) ListenAndServe(addr string) error {
	w.loadWallets()
	go w.startUpdatingCoinValue()

	r := mux.NewRouter()

	r.HandleFunc("/", w.Home)
	r.HandleFunc("/login", w.Login)
	r.HandleFunc("/gen", w.Generate)

	return http.ListenAndServe(addr, r)
}

func (w *Wallet) loadWallets() {
	ent, err := os.ReadDir("~/.bitcoin/wallets")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loading %d wallets\n", len(ent))
	for n, f := range ent {
		if f.IsDir() {
			err := w.bc.LoadWallet(f.Name())
			if err != nil {
				log.Printf("Couldn't load Wallet of %s: %s", f.Name(), err)
				continue
			}
			log.Printf("Loaded Wallet: %s (%d/%d)", f.Name(), n+1, len(ent))
		}
	}
}

func (w *Wallet) startUpdatingCoinValue() {
	go func() {
		for {
			w.cm.UpdateFiat()
			time.Sleep(time.Minute * 5)
		}
	}()
}

func (w *Wallet) redirectHome(wr http.ResponseWriter, r *http.Request) {
	http.Redirect(wr, r, fmt.Sprintf("http://%s", w.conf.Server.Address), 301)
}

func (w *Wallet) redirectLogin(wr http.ResponseWriter, r *http.Request) {
	http.Redirect(wr, r, fmt.Sprintf("http://%s/login", w.conf.Server.Address), 301)
}
