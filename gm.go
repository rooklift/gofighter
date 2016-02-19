package gofighter

import (
    "strconv"
)

type Level struct {
    Ok bool                 `json:"ok"`
    Error string            `json:"error,omitempty"`
    Account string          `json:"account"`
    InstanceId int          `json:"instanceId"`
    Venues []string         `json:"venues"`
    Tickers []string        `json:"tickers"`
    Balances map[string]int `json:"balances,omitempty"`
    BaseURL string          `json:"baseURL,omitempty"`       // These ain't part of the official response
    WebSocketURL string     `json:"websocketURL,omitempty"`  // but might be present in custom things...
}

const OFFICIAL_GM_URL = "https://www.stockfighter.io/gm"


func GMstart(info TradingInfo, levelname string)  (Level, error) {
    url := OFFICIAL_GM_URL + "/levels/" + levelname
    var ret Level
    err := get_json_from_url("POST", url, info.ApiKey, nil, &ret)
    return ret, err
}

func GMstop(info TradingInfo, instance int)  (Heartbeat, error) {
    url := OFFICIAL_GM_URL + "/instances/" + strconv.Itoa(instance) + "/stop"
    var ret Heartbeat
    err := get_json_from_url("POST", url, info.ApiKey, nil, &ret)
    return ret, err
}
