package gofighter

import (
    "fmt"
)

// The main Position type assumes it will be updated via its WebSocket

func (p * Position) Init(info TradingInfo)  {
    p.Tracker = make(chan Execution, 8192)          // Probably more than is needed
    go Tracker(info, p.Tracker)

    p.Info = info
}

func (p * Position) Update()  {

    if p.Tracker == nil {
        fmt.Println("WARNING: Attempt to use Position.Update() but channel was nil -- maybe Init() was never called?")
        return
    }

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
    return
}

func (p * Position) Print(currentprice int)  {
    if currentprice > 0 {
        nav := p.Cents + (p.Shares * currentprice)
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: $%d\n", p.Shares, p.Cents / 100, nav / 100)
    } else {
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: N/A\n", p.Shares, p.Cents / 100)
    }
}

// The SimplePosition type assumes it will be updated via reading cancels.
// It is the caller's responsibility to lock it if concurrency is happening.

func (p * SimplePosition) Init(info TradingInfo)  {
    p.Info = info
}

func (p * SimplePosition) UpdateFromMovement(move Movement)  {
    p.Cents += move.Cents
    p.Shares += move.Shares
}

func (p * SimplePosition) Print(currentprice int)  {
    if currentprice > 0 {
        nav := p.Cents + (p.Shares * currentprice)
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: $%d\n", p.Shares, p.Cents / 100, nav / 100)
    } else {
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: N/A\n", p.Shares, p.Cents / 100)
    }
}
