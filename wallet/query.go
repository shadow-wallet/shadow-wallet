package wallet

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

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
