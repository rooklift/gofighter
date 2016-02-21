package gofighter

// These are alternate, probably better ways of watching a market or position, completely thread safe.
// Call them as a new goroutine once and send a query via a channel when you want the info...

func MarketWatch(info TradingInfo, queries chan chan Market)  {

	m := Market{}	// We do all the initialisation here, no need to call .Init()
	m.Info = info
	m.LastPrice = -1
	m.Bid = -1
	m.Ask = -1

	ticker_channel := make(chan Quote, 256)	// Surely this is more than needed
	go Ticker(info, ticker_channel)

	for {
		select {

		case c := <- queries:

			c <- m

		case q := <- ticker_channel:

			saw_new_price := false

			// We store most of the quote in the Market info. Why not just store a quote? Well,
			// the quote has pointers which might be nil. Saving things as an int, with -1 as
			// a special "not present" value, makes life more convenient in the end. I hope.

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

			if q.BidSize != nil {
				m.BidSize = *q.BidSize
			} else {
				m.BidSize = -1
			}

			if q.AskSize != nil {
				m.AskSize = *q.AskSize
			} else {
				m.AskSize = -1
			}

			if q.BidDepth != nil {
				m.BidDepth = *q.BidDepth
			} else {
				m.BidDepth = -1
			}

			if q.AskDepth != nil {
				m.AskDepth = *q.AskDepth
			} else {
				m.AskDepth = -1
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

			if len(m.RecentPrices) > MARKET_PRICES_STORED {
				m.RecentPrices = m.RecentPrices[len(m.RecentPrices) - MARKET_PRICES_STORED :]
			}
		}
	}
}


func PositionWatch(info TradingInfo, queries chan chan Position)  {

	p := Position{}

	tracker_channel := make(chan Execution, 256)	// Surely this is more than needed
	go Tracker(info, tracker_channel)

    for {
        select {

		case c := <- queries:

			c <- p

        case msg := <- tracker_channel:

            if msg.Order.Direction == "buy" {
                p.Shares += msg.Filled
                p.Cents -= msg.Price * msg.Filled
            } else {
                p.Shares -= msg.Filled
                p.Cents += msg.Price * msg.Filled
            }
        }
    }
    return
}
