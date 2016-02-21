package gofighter

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

func GetPosition(query_channel chan chan Position)  Position {
    response_chan := make(chan Position)
    query_channel <- response_chan
    position := <- response_chan
    return position
}
