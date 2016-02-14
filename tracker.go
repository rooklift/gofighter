package gofighter

import (
	"bytes"
	"encoding/json"
)

func Tracker (ws_url string, account string, venue string, symbol string, results chan Execution) {

	url := ws_url + "/" + account + "/venues/" + venue + "/executions"
	conn := ws_connect_until_success(url)

	for {
		_, reader, err := conn.NextReader()
		if err != nil {

			conn.Close()
			conn = ws_connect_until_success(url)

		} else {

			var buf bytes.Buffer
			var e Execution

			buf.ReadFrom(reader)
			b := buf.Bytes()
			err = json.Unmarshal(b, &e)
			if err != nil {
				continue
			}
			results <- e
		}
	}
}
