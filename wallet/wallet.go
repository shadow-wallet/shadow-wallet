package wallet

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shadow-wallet/bitcoind"
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/wallet/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	loginTemplate, _ = template.New("assets/frontend/login.html", "assets/frontend/login.css")
)

type Wallet struct {
	bc *bitcoind.Bitcoind
	cm *coinmarket.CoinMarket
}

func New(conf Config) (*Wallet, error) {
	bc, err := bitcoind.New(conf.Daemon.Address,
		conf.Daemon.Username, conf.Daemon.Password)
	if err != nil {
		return nil, err
	}
	w := &Wallet{
		cm: coinmarket.New(conf.CoinMarket.APIKEY),
		bc: bc,
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

func queryToJSON(body []byte) []byte {
	if len(body) == 0 {
		return body
	}
	j := map[string]any{}
	elm := strings.Split(string(body), "&")
	for _, e := range elm {
		spl := strings.Split(e, "=")
		j[spl[0]] = spl[1]
	}
	dat, err := json.Marshal(j)
	if err != nil {
		log.Fatalln(err)
	}
	return dat
}

func (w *Wallet) redirectHome(wr http.ResponseWriter, r *http.Request) {
	http.Redirect(wr, r, "http://127.0.0.1:8080", 301)
}

func (w *Wallet) redirectLogin(wr http.ResponseWriter, r *http.Request) {
	http.Redirect(wr, r, "http://127.0.0.1:8080/login", 301)
}

func unmarshalQuery(body io.ReadCloser, dst any) ([]byte, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return b, err
	}
	if len(b) == 0 {
		return b, nil
	}
	buf := queryToJSON(b)
	return buf, json.Unmarshal(buf, &dst)
}
