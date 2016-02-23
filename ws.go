package gofighter

import (
    "bytes"
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
)

func ws_connect_until_success(url string)  (*websocket.Conn) {

    var dialer websocket.Dialer
    var header http.Header

    var conn * websocket.Conn
    var err error

    for {
        conn, _, err = dialer.Dial(url, header)
        if err != nil {
            time.Sleep(2 * time.Second)
            continue
        }
        break
    }

    return conn
}

func Tracker(info TradingInfo, results chan Execution)  {

    url := info.WebSocketURL + "/" + info.Account + "/venues/" + info.Venue + "/executions"
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
            if err != nil || e.Ok == false {
                continue
            }
            results <- e
        }
    }
}

func Ticker(info TradingInfo, results chan Quote)  {

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
            if err != nil || q.Ok == false {        // Note that q is the TickerQuote wrapper and q.Ok is not a pointer
                continue
            }
            results <- q.RawQuote.Quote()           // Convert the full-of-pointers struct to a normal Quote
                                                    // (it might lack the Ok field though, server don't send that)
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
