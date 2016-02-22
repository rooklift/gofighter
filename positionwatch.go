package gofighter

// The watchers are the recommended way to use WebSockets to watch a market
// or position, completely thread-safe. Call them as a new goroutine once and
// send a query via a channel when you want the info... that way you get a
// copy of the info via the channel, and can safely read from it.

// Having said that, using WebSockets to monitor one's position might be
// subject to delay, so updating a Position directly from cancel results is
// also a plausible strategy.

func PositionWatch(info TradingInfo, queries chan chan Position)  {

    p := Position{}
    p.Info = info

    tracker_channel := make(chan Execution, 256)    // Surely this is more than needed
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

func GetPosition(query_channel chan chan Position)  Position {
    response_chan := make(chan Position)
    query_channel <- response_chan
    position := <- response_chan
    return position
}
