package gofighter

import (
    "encoding/json"
    "sync"
)

type Execution struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Account             string      `json:"account"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Order               Order       `json:"order"`
    StandingId          int         `json:"standingId"`
    IncomingId          int         `json:"incomingId"`
    Price               int         `json:"price"`
    Filled              int         `json:"filled"`
    FilledAt            string      `json:"filledAt"`
    StandingComplete    bool        `json:"standingComplete"`
    IncomingComplete    bool        `json:"incomingComplete"`
}

type Heartbeat struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error"`
}

type VenueHeartbeat struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Venue               string      `json:"venue"`
}

type Venue struct {
    Id                  int         `json:"id"`
    Name                string      `json:"name"`
    Venue               string      `json:"venue"`
    State               string      `json:"state"`
}

type VenueList struct {
    Ok                  bool        `json:"id"`                 // Bug on official server
    Error               string      `json:"error,omitempty"`
    Venues              []Venue     `json:"venues"`
}

type Symbol struct {
    Name                string      `json:"name"`
    Symbol              string      `json:"symbol"`
}

type StockList struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Symbols             []Symbol    `json:"symbols"`
}

type BookEntry struct {
    Price               int         `json:"price"`
    Qty                 int         `json:"qty"`
    IsBuy               bool        `json:"isBuy"`
}

type OrderBook struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Ts                  string      `json:"ts"`
    Bids                []BookEntry `json:"bids"`
    Asks                []BookEntry `json:"asks"`
}

type Fill struct {
    Price               int         `json:"price"`
    Qty                 int         `json:"qty"`
    Ts                  string      `json:"ts"`
}

type Order struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Account             string      `json:"account"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Direction           string      `json:"direction"`
    OrderType           string      `json:"orderType"`
    Ts                  string      `json:"ts"`
    OriginalQty         int         `json:"originalQty"`
    Qty                 int         `json:"qty"`
    Price               int         `json:"price"`
    TotalFilled         int         `json:"totalFilled"`
    Id                  int         `json:"id"`
    Fills               []Fill      `json:"fills"`
    Open                bool        `json:"open"`
}

type OrderList struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Orders              []Order     `json:"orders"`
}

// The following things are created purely or mostly client-side
// (not marshalled directly from the server)...

type RawOrder struct {      // This is what actually gets marshalled and sent to the server
    Account             string      `json:"account"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Direction           string      `json:"direction"`
    OrderType           string      `json:"orderType"`
    Qty                 int         `json:"qty"`
    Price               int         `json:"price"`
}

type ShortOrder struct {    // Like RawOrder, but missing everything we can get from TradingInfo
    Direction           string      `json:"direction"`
    OrderType           string      `json:"orderType"`
    Qty                 int         `json:"qty"`
    Price               int         `json:"price"`
}

type TradingInfo struct {
    BaseURL             string      `json:"baseURL"`
    WebSocketURL        string      `json:"websocketURL,omitempty"`
    ApiKey              string      `json:"apiKey"`
    Account             string      `json:"account"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
}

type Movement struct {
    Cents               int
    Shares              int
}

type Position struct {
    Info                TradingInfo
    Cents               int
    Shares              int
    Lock                sync.Mutex
}

type Market struct {
    Info                TradingInfo
    RecentPrices        []int
    Quote               Quote
    Lock                sync.Mutex
}

// The following methods and interface allow either a RawOrder or a ShortOrder to be passed to base.Execute()

func (original RawOrder) MakeShortOrder() ShortOrder  {
    return ShortOrder{
        Direction: original.Direction,
        OrderType: original.OrderType,
        Qty: original.Qty,
        Price: original.Price,
    }
}

func (original ShortOrder) MakeShortOrder() ShortOrder  {
    return original
}

type ShortOrderer interface {
    MakeShortOrder() ShortOrder
}

// Quotes are special due to the possibility of important fields being absent...

type Quote struct {

    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Bid                 int         `json:"bid,omitempty"`                  // Gets set to -1 if absent
    Ask                 int         `json:"ask,omitempty"`                  // Gets set to -1 if absent
    BidSize             int         `json:"bidSize"`
    AskSize             int         `json:"askSize"`
    BidDepth            int         `json:"bidDepth"`
    AskDepth            int         `json:"askDepth"`
    Last                int         `json:"last,omitempty"`                 // Gets set to -1 if absent
    LastSize            int         `json:"lastSize,omitempty"`
    LastTrade           string      `json:"lastTrade,omitempty"`
    QuoteTime           string      `json:"quoteTime"`
}

// The Quote is special since I want some absent values to get set to -1
// instead of 0. Therefore, a custom UnmarshalJSON() method is required.
// The Go unmarshaller calls this method automatically as needed...

func (q * Quote) UnmarshalJSON(b []byte)  error {

    // First unmarshal quotes into pointers so we can tell the
    // difference between field-not-present and actually-zero values.

    type RawQuote struct {
        Ok                  *bool       `json:"ok"`
        Error               *string     `json:"error"`
        Venue               *string     `json:"venue"`
        Symbol              *string     `json:"symbol"`
        Bid                 *int        `json:"bid"`                // Can be absent
        Ask                 *int        `json:"ask"`                // Can be absent
        BidSize             *int        `json:"bidSize"`
        AskSize             *int        `json:"askSize"`
        BidDepth            *int        `json:"bidDepth"`
        AskDepth            *int        `json:"askDepth"`
        Last                *int        `json:"last"`               // Can be absent
        LastSize            *int        `json:"lastSize"`           // Can be absent
        LastTrade           *string     `json:"lastTrade"`          // Can be absent
        QuoteTime           *string     `json:"quoteTime"`
    }

    r := RawQuote{}
    err := json.Unmarshal(b, &r)

    if r.Ok         != nil { q.Ok           = *r.Ok         }
    if r.Error      != nil { q.Error        = *r.Error      }
    if r.Venue      != nil { q.Venue        = *r.Venue      }
    if r.Symbol     != nil { q.Symbol       = *r.Symbol     }
    if r.BidSize    != nil { q.BidSize      = *r.BidSize    }
    if r.AskSize    != nil { q.AskSize      = *r.AskSize    }
    if r.BidDepth   != nil { q.BidDepth     = *r.BidDepth   }
    if r.AskDepth   != nil { q.AskDepth     = *r.AskDepth   }
    if r.LastSize   != nil { q.LastSize     = *r.LastSize   }
    if r.LastTrade  != nil { q.LastTrade    = *r.LastTrade  }
    if r.QuoteTime  != nil { q.QuoteTime    = *r.QuoteTime  }

    // If they are nil, zero values for the above are acceptable.
    // But we must distinguish between 0-price and no-price:

    if r.Bid != nil {
        q.Bid = *r.Bid
    } else {
        q.Bid = -1
    }

    if r.Ask != nil {
        q.Ask = *r.Ask
    } else {
        q.Ask = -1
    }

    if r.Last != nil {
        q.Last = *r.Last
    } else {
        q.Last = -1
    }

    return err
}

// The Ticker WebSocket sends the following thing for some reason:

type TickerQuote struct {
    Ok                  bool        `json:"ok"`
    Error               string      `json:"error,omitempty"`
    Quote               Quote       `json:"quote"`
}
