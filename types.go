package gofighter

// Use pointers so missing values are nil after marshalling
// (this is especially important for a quote).

// Also, I assume without evidence that everything that has an "ok"
// field will produce an "error" field if "ok" is false.

type Quote struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Venue       *string     `json:"venue"`
    Symbol      *string     `json:"symbol"`
    Bid         *int        `json:"bid"`            // Can be absent
    Ask         *int        `json:"ask"`            // Can be absent
    BidSize     *int        `json:"bidSize"`
    AskSize     *int        `json:"askSize"`
    BidDepth    *int        `json:"bidDepth"`
    AskDepth    *int        `json:"askDepth"`
    Last        *int        `json:"last"`           // Can be absent
    LastSize    *int        `json:"lastSize"`       // Can be absent
    LastTrade   *string     `json:"lastTrade"`      // Can be absent
    QuoteTime   *string     `json:"quoteTime"`
}

type TickerQuote struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Quote       *Quote      `json:"quote"`
}

type Heartbeat struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
}

type VenueHeartbeat struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Venue       *string     `json:"venue"`
}

type Venue struct {
    Id          *int        `json:"id"`
    Name        *string     `json:"name"`
    Venue       *string     `json:"venue"`
    State       *string     `json:"state"`
}

type VenueList struct {
    Ok          *bool       `json:"id"`             // Bug on official server, will presumably be fixed at some point
    Error       *string     `json:"error"`
    Venues      []Venue     `json:"venues"`
}

type Symbol struct {
    Name        *string     `json:"name"`
    Symbol      *string     `json:"symbol"`
}

type StockList struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Symbols     []Symbol    `json:"symbols"`
}

type BookEntry struct {
    Price       *int        `json:"price"`
    Qty         *int        `json:"qty"`
    IsBuy       *int        `json:"isBuy"`
}

type OrderBook struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Venue       *string     `json:"venue"`
    Symbol      *string     `json:"symbol"`
    Ts          *string     `json:"ts"`
    Bids        []BookEntry `json:"bids"`
    Asks        []BookEntry `json:"asks"`
}

type Fill struct {
    Price       *int        `json:"price"`
    Qty         *int        `json:"qty"`
    Ts          *string     `json:"ts"`
}

type Order struct {
    Ok          *bool       `json:"ok"`
    Error       *string     `json:"error"`
    Account     *string     `json:"account"`
    Venue       *string     `json:"venue"`
    Symbol      *string     `json:"symbol"`
    Direction   *string     `json:"direction"`
    OrderType   *string     `json:"orderType"`
    Ts          *string     `json:"ts"`
    OriginalQty *int        `json:"originalQty"`
    Qty         *int        `json:"qty"`
    Price       *int        `json:"price"`
    TotalFilled *int        `json:"totalFilled"`
    Id          *int        `json:"id"`
    Fills       []Fill      `json:"fills"`
    Open        *bool       `json:"open"`
}

// These are created by the client to be sent to the server,
// so more convenient not to make them pointers...
type RawOrder struct {
    Account     string      `json:"account"`
    Venue       string      `json:"venue"`
    Symbol      string      `json:"symbol"`
    Direction   string      `json:"direction"`
    OrderType   string      `json:"orderType"`
    Qty         int         `json:"qty"`
    Price       int         `json:"price"`
}
