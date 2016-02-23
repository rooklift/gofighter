package gofighter

import (
    "fmt"
)

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
    if m.Quote.Bid == -1 || m.Quote.Ask == -1 {
        return -1
    }
    return m.Quote.Ask - m.Quote.Bid
}

func (m * Market) Print()  {
    if m.Quote.Last == -1 {
        fmt.Println("No market activity yet, or server down")
        return
    }

    var bidstr string
    if m.Quote.Bid != -1 {
        bidstr = fmt.Sprintf("%5d", m.Quote.Bid)
    } else {
        bidstr = ""
    }

    var askstr string
    if m.Quote.Ask != -1 {
        askstr = fmt.Sprintf("%-5d", m.Quote.Ask)
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
