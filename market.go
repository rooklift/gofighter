package gofighter

import (
    "fmt"
)

const MARKET_PRICES_STORED = 200

func (m * Market) Init(info TradingInfo, tickerfunction func(TradingInfo, chan Quote))  {
    m.Ticker = make(chan Quote, 8192)           // Probably more than is needed
    go tickerfunction(info, m.Ticker)

    m.Info = info
    m.LastPrice = -1
    m.Bid = -1
    m.Ask = -1
}

func (m * Market) Update()  int {

    // Update the market from the WebSocket results channel.
    // Assumes m.Init() has been called (once ever).
    // Return the number of WebSocket messages read.

    var count int

    loop:
    for {
        select {

            case q := <- m.Ticker:

                count++

                saw_new_price := false

                if q.Bid != nil {
                    m.Bid = *q.Bid
                } else {
                    m.Bid = -1
                }

                if q.Ask != nil {
                    m.Ask = *q.Ask
                } else {
                    m.Ask = -1
                }

                if q.Last != nil {
                    if m.LastPrice != *q.Last {
                        saw_new_price = true
                    }
                    m.LastPrice = *q.Last
                } else {
                    m.LastPrice = -1
                }

                if q.LastTrade != nil {
                    if m.LastTime != *q.LastTrade {
                        saw_new_price = true
                    }
                    m.LastTime = *q.LastTrade
                } else {
                    m.LastTime = ""
                }

                if saw_new_price {
                    m.RecentPrices = append(m.RecentPrices, m.LastPrice)
                }

            default:
                break loop
        }
    }

    if len(m.RecentPrices) > MARKET_PRICES_STORED {
        m.RecentPrices = m.RecentPrices[len(m.RecentPrices) - MARKET_PRICES_STORED :]
    }

    return count
}

func (m * Market) Average()  int {
    if len(m.RecentPrices) == 0 {
        return -1
    }

    var total int64
    var count int
    var price int

    for count, price = range m.RecentPrices {
        total += int64(price)
    }

    return int(total / int64(count + 1))        // Adjustment required for zeroth indexing
}

func (m * Market) AverageSansOutliers()  int {
    av := m.Average()
    if av == -1 {
        return -1
    }

    var total int64
    var count int

    lower_bound := (av * 10) / 12
    upper_bound := (av * 12) / 10

    for _, price := range m.RecentPrices {
        if price > lower_bound && price < upper_bound {
            total += int64(price)
            count += 1
        }
    }

    if count == 0 {
        return -1
    } else {
        return int(total / int64(count))
    }
}

func (m * Market) Spread()  int {
    if m.Bid == -1 || m.Ask == -1 {
        return -1
    }
    return m.Ask - m.Bid
}

func (m * Market) Print()  {
    if m.LastPrice == -1 {
        fmt.Println("No market activity yet, or server down")
        return
    }

    var bidstr string
    if m.Bid != -1 {
        bidstr = fmt.Sprintf("%5d", m.Bid)
    } else {
        bidstr = ""
    }

    var askstr string
    if m.Ask != -1 {
        askstr = fmt.Sprintf("%-5d", m.Ask)
    } else {
        askstr = ""
    }

    var arrow string
    if m.Spread() != -1 {
        arrow = fmt.Sprintf("<--- %3d --->", m.Spread())
    } else {
        arrow = "<----------->"
    }

    fmt.Printf("%5s %s %5s ... averages: %5d (n == %3d) |> %-5d\n",
               bidstr, arrow, askstr, m.Average(), len(m.RecentPrices), m.AverageSansOutliers())
    return
}
