package gofighter

import (
	"bytes"
	"encoding/json"
)

func Tracker(ws_url string, account string, venue string, symbol string, results chan Execution) {

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

func PositionUpdater(ws_url string, account string, venue string, symbol string,
					 pos * Position, ws_results chan Execution, updated_pos_chan chan Position) {

	// The position updater does 3 things:
	//     * Updates the Position, which can optionally be in shared memory (otherwise, it creates one if the ptr is nil)
	//     * Optionally relays tracker messages to a channel
	//     * Optionally sends current position messages to a channel

	if pos == nil {
		pos = new(Position)
	}

	pos.Lock.Lock()
	pos.Account = account
	pos.Venue = venue
	pos.Symbol = symbol
	pos.Lock.Unlock()

	local_chan := make(chan Execution, 64)
	go Tracker(ws_url, account, venue, symbol, local_chan)

	var updatedpos Position

	for {
		msg := <- local_chan

		pos.Lock.Lock()
		if msg.Order.Direction == "buy" {
			pos.Shares += msg.Filled
			pos.Cents -= msg.Price * msg.Filled
		} else {
			pos.Shares -= msg.Filled
			pos.Cents += msg.Price * msg.Filled
		}
		updatedpos = *pos				// Set this inside the lock but send it later (outside the lock), might help avoid deadlocks
		pos.Lock.Unlock()

		if ws_results != nil {			// Optionally pass on the raw WS msg from Tracker
			ws_results <- msg
		}

		if updated_pos_chan != nil {	// Optionally send the current value of pos
			updated_pos_chan <- updatedpos
		}
	}
}
