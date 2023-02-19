package app

import (
	"encoding/json"
	"fmt"
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/app/auth"
	"log"
	"net/http"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	username, ok := auth.Authenticated(r)
	if !ok {
		redirectLogin(w, r)
		return
	}

	var dst any
	body, err := unmarshalQuery(r.Body, &dst)
	if err != nil {
		log.Fatalln(err)
	}

	if actions, ok := dst.(map[string]any); ok && a.handleAction(actions, w, username, body) {
		return
	}

	b, err := a.bc.GetBalance(username, 1)
	if err != nil {
		_, _ = fmt.Fprintf(w, "error getting wallet information, issue reported to staff")
		return
	}
	_, _ = fmt.Fprintf(w, "LTC:  CA$%.2f\nDOGE: CA$%.2f\nBTC:  CA$%.2f\n\nBal: CA$%.2f (%f BTC)\n", coinmarket.LTC().Price(), coinmarket.DOGE().Price(), coinmarket.BTC().Price(), b*coinmarket.BTC().Price(), b)
}

func (a *App) handleAction(actions map[string]any, w http.ResponseWriter, username string, body []byte) bool {
	switch actions["action"] {
	case "send":
		a.send(w, username, body)
		return true
	}
	return false
}

type transactionInfo struct {
	Amount    float64 `json:"amount"`
	Recipient string  `json:"recipient"`
}

func (a *App) send(w http.ResponseWriter, from string, body []byte) {
	var txInfo transactionInfo
	_ = json.Unmarshal(body, &txInfo)

	txID, err := a.bc.SendToAddress(from, txInfo.Recipient, txInfo.Amount, "", "", true)
	if err != nil {
		_, _ = fmt.Fprintf(w, "error sending funds, issue reported to staff")
		return
	}
	_, _ = fmt.Fprintf(w, "Transaction ID: %s", txID)
}
