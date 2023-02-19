package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shadow-wallet/bitcoind"
	"github.com/shadow-wallet/shadow-wallet/app/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	loginTemplate, _ = template.New("assets/frontend/login.html", "assets/frontend/login.css")
)

type App struct {
	bc *bitcoind.Bitcoind
}

func NewApp(bc *bitcoind.Bitcoind) *App {
	return &App{
		bc: bc,
	}
}

func (a *App) ListenAndServe(addr string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", a.Home)
	r.HandleFunc("/login", a.Login)
	r.HandleFunc("/gen", a.Generate)

	return http.ListenAndServe(addr, r)
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

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://127.0.0.1:8080", 301)
}

func redirectLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://127.0.0.1:8080/login", 301)
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
