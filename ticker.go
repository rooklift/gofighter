package gofighter

import (
	"bytes"
	"encoding/json"
)

func Ticker (ws_url string, account string, venue string, symbol string, results chan Quote) {

	conn := ws_connect_until_success(ws_url + "/" + account + "/venues/" + venue + "/tickertape")

	for {
		_, reader, err := conn.NextReader()
		if err != nil {

			conn.Close()
			conn = ws_connect_until_success(ws_url + "/" + account + "/venues/" + venue + "/tickertape")

		} else {

			var buf bytes.Buffer
			var q TickerQuote

			buf.ReadFrom(reader)
			s := buf.String()
			err = json.Unmarshal([]byte(s), &q)
			if err != nil {
				continue
			}
			results <- *q.Quote
		}
	}
}
