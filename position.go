package gofighter

func (p * Position) Init(info TradingInfo)  {
    p.Tracker = make(chan Execution, 8192)          // Probably more than is needed
    go Tracker(info, p.Tracker)
}

func (p * Position) Update()  {

    loop:
    for {
        select {

        case msg := <- p.Tracker:

            if msg.Order.Direction == "buy" {
                p.Shares += msg.Filled
                p.Cents -= msg.Price * msg.Filled
            } else {
                p.Shares -= msg.Filled
                p.Cents += msg.Price * msg.Filled
            }

        default:
            break loop
        }
    }
}
