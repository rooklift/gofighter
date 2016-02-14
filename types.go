package gofighter

// Use pointers so missing values are nil after marshalling
// (this is especially important for a quote). This also means
// that the JSON printer can skip nil fields if the omitempty
// tag is given (it WILL still print the field if it's a valid
// pointer to a zero value -- i.e. (int) 0 or (string) ""

// Also, I assume without evidence that everything that has an
// "ok" field will produce an "error" field if "ok" is false.

type Quote struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`                // Usually absent
    Venue               *string     `json:"venue"`
    Symbol              *string     `json:"symbol"`
    Bid                 *int        `json:"bid,omitempty"`                  // Can be absent, still prints if 0
    Ask                 *int        `json:"ask,omitempty"`                  // Can be absent, still prints if 0
    BidSize             *int        `json:"bidSize"`
    AskSize             *int        `json:"askSize"`
    BidDepth            *int        `json:"bidDepth"`
    AskDepth            *int        `json:"askDepth"`
    Last                *int        `json:"last,omitempty"`                 // Can be absent, still prints if 0
    LastSize            *int        `json:"lastSize,omitempty"`             // Can be absent
    LastTrade           *string     `json:"lastTrade,omitempty"`            // Can be absent
    QuoteTime           *string     `json:"quoteTime"`
}

type TickerQuote struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Quote               *Quote      `json:"quote"`
}

type Execution struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Account             *string     `json:"account"`
    Venue               *string     `json:"venue"`
    Symbol              *string     `json:"symbol"`
    Order               *Order      `json:"order"`
    StandingId          *int        `json:"standingId"`
    IncomingId          *int        `json:"incomingId"`
    Price               *int        `json:"price"`
    Filled              *int        `json:"filled"`
    FilledAt            *string     `json:"filledAt"`
    StandingComplete    *bool       `json:"standingComplete"`
    IncomingComplete    *bool       `json:"incomingComplete"`
}

type Heartbeat struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error"`
}

type VenueHeartbeat struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Venue               *string     `json:"venue"`
}

type Venue struct {
    Id                  *int        `json:"id"`
    Name                *string     `json:"name"`
    Venue               *string     `json:"venue"`
    State               *string     `json:"state"`
}

type VenueList struct {
    Ok                  *bool       `json:"id"`                 // Bug on official server
    Error               *string     `json:"error,omitempty"`
    Venues              []Venue     `json:"venues"`
}

type Symbol struct {
    Name                *string     `json:"name"`
    Symbol              *string     `json:"symbol"`
}

type StockList struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Symbols             []Symbol    `json:"symbols"`
}

type BookEntry struct {
    Price               *int        `json:"price"`
    Qty                 *int        `json:"qty"`
    IsBuy               *bool       `json:"isBuy"`
}

type OrderBook struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Venue               *string     `json:"venue"`
    Symbol              *string     `json:"symbol"`
    Ts                  *string     `json:"ts"`
    Bids                []BookEntry `json:"bids"`
    Asks                []BookEntry `json:"asks"`
}

type Fill struct {
    Price               *int        `json:"price"`
    Qty                 *int        `json:"qty"`
    Ts                  *string     `json:"ts"`
}

type Order struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Account             *string     `json:"account"`
    Venue               *string     `json:"venue"`
    Symbol              *string     `json:"symbol"`
    Direction           *string     `json:"direction"`
    OrderType           *string     `json:"orderType"`
    Ts                  *string     `json:"ts"`
    OriginalQty         *int        `json:"originalQty"`
    Qty                 *int        `json:"qty"`
    Price               *int        `json:"price"`
    TotalFilled         *int        `json:"totalFilled"`
    Id                  *int        `json:"id"`
    Fills               []Fill      `json:"fills"`
    Open                *bool       `json:"open"`
}

type OrderList struct {
    Ok                  *bool       `json:"ok"`
    Error               *string     `json:"error,omitempty"`
    Orders              []Order     `json:"orders"`
}

// These are created by the client to be sent to the server,
// so more convenient not to make them pointers...
type RawOrder struct {
    Account             string      `json:"account"`
    Venue               string      `json:"venue"`
    Symbol              string      `json:"symbol"`
    Direction           string      `json:"direction"`
    OrderType           string      `json:"orderType"`
    Qty                 int         `json:"qty"`
    Price               int         `json:"price"`
}
