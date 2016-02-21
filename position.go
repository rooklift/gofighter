package gofighter

import (
    "fmt"
)

// If a Position is in shared memory, it is the caller's responsibility to ensure safety.
// Although using the PositionWatch() function is nice, it might be too delayed due to
// WebSocket issues, so updating directly may be needed...

func (p * Position) Print(currentprice int)  {
    if currentprice > 0 {
        nav := p.Cents + (p.Shares * currentprice)
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: $%d\n", p.Shares, p.Cents / 100, nav / 100)
    } else {
        fmt.Printf("Shares: %d, Dollars: $%d, NAV: N/A\n", p.Shares, p.Cents / 100)
    }
}

func (p * Position) UpdateFromMovement(move Movement)  {
    p.Cents += move.Cents
    p.Shares += move.Shares
}
