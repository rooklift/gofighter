package gofighter

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
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
    BaseURL string          `json:"baseURL,omitempty"`       // These ain't part of the
    WebSocketURL string     `json:"websocketURL,omitempty"`  // official response but might
    ApiKey string           `json:"apiKey,omitempty"`        // be present in custom things
}

const OFFICIAL_BASE_URL = "https://api.stockfighter.io/ob/api"
const OFFICIAL_WS_URL = "https://api.stockfighter.io/ob/api/ws"
const OFFICIAL_GM_URL = "https://www.stockfighter.io/gm"

const GM_DIRECTORY = "gm"


func GMstart(apikey string, levelname string)  (Level, error) {

    var err error
    if apikey == "" {
        apikey, err = LoadAPIKey("api_key.txt")
        if err != nil {
            fmt.Println(err)
        }
    }

    url := OFFICIAL_GM_URL + "/levels/" + levelname
    var ret Level
    err = get_json_from_url("POST", url, apikey, nil, &ret)
    return ret, err
}

func GMstop(apikey string, instance int)  (Heartbeat, error) {

    var err error
    if apikey == "" {
        apikey, err = LoadAPIKey("api_key.txt")
        if err != nil {
            fmt.Println(err)
        }
    }

    url := OFFICIAL_GM_URL + "/instances/" + strconv.Itoa(instance) + "/stop"
    var ret Heartbeat
    err = get_json_from_url("POST", url, apikey, nil, &ret)
    return ret, err
}

func LoadGMfile(levelname string)  (Level, error) {
    filename := GM_DIRECTORY + "/" + levelname + "_info.txt"

    var ret Level

    s, err := ioutil.ReadFile(filename)
    if err != nil {
        ret.Error = err.Error()
        return ret, err
    }

    err = json.Unmarshal(s, &ret)
    if err != nil {
        ret.Error = err.Error()
        return ret, err
    }

    return ret, nil
}

func SaveGMfile(levelname string, level Level)  {

    if levelname == "" {
        fmt.Println(`SaveGMfile() : filename was ""`)
        return
    }

    filename := GM_DIRECTORY + "/" + levelname + "_info.txt"
    os.Mkdir(GM_DIRECTORY, 0777)

    s, _ := MakeJSON(level)
    err := ioutil.WriteFile(filename, []byte(s), 0777)
    if err != nil {
        fmt.Println(err)
    }
    return
}

func TradingInfoFromLevel(level Level)  TradingInfo {

    info := TradingInfo{}

    info.BaseURL = level.BaseURL
    info.WebSocketURL = level.WebSocketURL
    info.ApiKey = level.ApiKey
    info.Account = level.Account
    if len(level.Venues) > 0 {
        info.Venue = level.Venues[0]
    }
    if len(level.Tickers) > 0 {
        info.Symbol = level.Tickers[0]
    }

    if info.BaseURL == "" {info.BaseURL = OFFICIAL_BASE_URL}
    if info.WebSocketURL == "" {info.WebSocketURL = OFFICIAL_WS_URL}
    if info.ApiKey == "" {
        var err error
        info.ApiKey, err = LoadAPIKey("api_key.txt")
        if err != nil {
            fmt.Println(err)
        }
    }

    return info
}

func TradingInfoFromName(levelname string)  TradingInfo {

    // Assuming we have already contacted the GM and written a GM info file,
    // this function returns a TradingInfo struct, using sensible default
    // values for the URLs and API key.

    level, err := LoadGMfile(levelname)
    if err != nil {
        fmt.Println(err)
    }
    return TradingInfoFromLevel(level)
}

func NameFromUser()  string {
    var known = make(map[string]string)

    s, err := ioutil.ReadFile("known_levels.json")
    if err != nil {
        fmt.Println(err)
    }

    json.Unmarshal(s, &known)
    PrintJSON(known)

    fmt.Println("\nEnter level number...")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    return known[scanner.Text()]
}
