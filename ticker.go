package gofighter

import (
	"bytes"
	"encoding/json"
	"time"
)

func Ticker(info TradingInfo, results chan Quote) {

	url := info.WebSocketURL + "/" + info.Account + "/venues/" + info.Venue + "/tickertape"
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

func FakeTicker(info TradingInfo, results chan Quote)  {

    // Poor man's tickertape without WebSockets...

    for {
        res, err := GetQuote(info)
        if err != nil {
            continue
        }
        results <- res
        time.Sleep(500 * time.Millisecond)
    }
}
