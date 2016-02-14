package gofighter

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func ws_connect_until_success (url string) (*websocket.Conn) {

	var dialer websocket.Dialer
	var header http.Header

	var conn * websocket.Conn
	var err error

	for {
		conn, _, err = dialer.Dial(url, header)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	return conn
}
