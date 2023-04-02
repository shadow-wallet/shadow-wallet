package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/shadow-wallet/coinmarket"
	"github.com/shadow-wallet/shadow-wallet/wallet/auth"
	"log"
	"net/http"
)

func (w *Wallet) Home(wr http.ResponseWriter, r *http.Request) {
	username, ok := auth.Authenticated(r)
	if !ok {
		w.redirectLogin(wr, r)
		return
	}

	var dst any
	body, err := unmarshalQuery(r.Body, &dst)
	if err != nil {
		log.Fatalln(err)
	}

	if actions, ok := dst.(map[string]any); ok && w.handleAction(actions, wr, username, body) {
		return
	}

	b, err := w.bc.GetBalance(username, 1)
	if err != nil {
		_, _ = fmt.Fprintf(wr, "error getting Wallet information, issue reported to staff")
		return
	}
	_, _ = fmt.Fprintf(wr, "LTC:  CA$%.2f\nDOGE: CA$%.2f\nBTC:  CA$%.2f\n\nBal: CA$%.2f (%f BTC)\n", coinmarket.LTC().Price(), coinmarket.DOGE().Price(), coinmarket.BTC().Price(), b*coinmarket.BTC().Price(), b)
}

func (w *Wallet) handleAction(actions map[string]any, wr http.ResponseWriter, username string, body []byte) bool {
	switch actions["action"] {
	case "send":
		w.send(wr, username, body)
		return true
	}
	return false
}

type transactionInfo struct {
	Amount    float64 `json:"amount"`
	Recipient string  `json:"recipient"`
}

func (w *Wallet) send(wr http.ResponseWriter, from string, body []byte) {
	var txInfo transactionInfo
	_ = json.Unmarshal(body, &txInfo)

	txID, err := w.bc.SendToAddress(from, txInfo.Recipient, txInfo.Amount, "", "", true)
	if err != nil {
		_, _ = fmt.Fprintf(wr, "error sending funds, issue reported to staff")
		return
	}
	_, _ = fmt.Fprintf(wr, "Transaction ID: %s", txID)
}
