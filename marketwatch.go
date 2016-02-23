package gofighter

// The watchers are the recommended way to use WebSockets to watch a market
// or position, completely thread-safe. Call them as a new goroutine once and
// send a query via a channel when you want the info... that way you get a
// copy of the info via the channel, and can safely read from it.

const MARKET_PRICES_STORED = 200

func MarketWatch(info TradingInfo, queries chan chan Market)  {

    m := Market{}
    m.Info = info
    m.Quote = Quote{    // This is what the quote would look like before any trades
        Ok: true,
        Error: "",
        Venue: info.Venue,
        Symbol: info.Symbol,
        Bid: -1,                // Special value we use for "not present in JSON"
        Ask: -1,                // Special value we use for "not present in JSON"
        BidSize: 0,
        AskSize: 0,
        BidDepth: 0,
        AskDepth: 0,
        Last: -1,               // Special value we use for "not present in JSON"
        LastSize: 0,
        LastTrade: "",
        QuoteTime: "",
    }

    ticker_channel := make(chan Quote, 256)        // Surely this is more than needed
    go Ticker(info, ticker_channel)

    for {
        select {

        case c := <- queries:

            c <- m

        case q := <- ticker_channel:

            saw_new_price := false

            if m.Quote.Last != q.Last {
                saw_new_price = true
            }

            if m.Quote.LastTrade != q.LastTrade {
                saw_new_price = true
            }

            m.Quote = q     // Do this after the above, of course

            if saw_new_price {
                m.RecentPrices = append(m.RecentPrices, m.Quote.Last)
            }

            if len(m.RecentPrices) > MARKET_PRICES_STORED {
                m.RecentPrices = m.RecentPrices[len(m.RecentPrices) - MARKET_PRICES_STORED :]
            }
        }
    }
}

func GetMarket(query_channel chan chan Market)  Market {
    response_chan := make(chan Market)
    query_channel <- response_chan
    market := <- response_chan
    return market
}
