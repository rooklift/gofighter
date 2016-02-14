package gofighter

import (
	"bytes"
	"encoding/json"
)

func Ticker (ws_url string, account string, venue string, symbol string, results chan Quote) {

	url := ws_url + "/" + account + "/venues/" + venue + "/tickertape"
	conn := ws_connect_until_success(url)

	for {
		_, reader, err := conn.NextReader()
		if err != nil {

			conn.Close()
			conn = ws_connect_until_success(url)

		} else {

			var buf bytes.Buffer
			var q TickerQuote

			buf.ReadFrom(reader)
			b := buf.Bytes()
			err = json.Unmarshal(b, &q)
			if err != nil {
				continue
			}
			results <- q.Quote
		}
	}
}
