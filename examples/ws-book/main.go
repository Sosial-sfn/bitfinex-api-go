package main

import (
    "fmt"
    "log"

    "github.com/bitfinexcom/bitfinex-api-go"
)

func main() {
    c := bitfinex.NewClient()

    // in case your proxy is using a non valid certificate set to TRUE
    c.WebSocketTLSSkipVerify = false

    err := c.WebSocket.Connect()
    if err != nil {
        log.Fatal("Error connecting to web socket : ", err)
    }
    defer c.WebSocket.Close()

    book_btcusd_chan := make(chan []float64)
    book_ltcusd_chan := make(chan []float64)
    trades_chan := make(chan []float64)
    ticker_chan := make(chan []float64)

    c.WebSocket.AddSubscribe(bitfinex.CHAN_BOOK, bitfinex.BTCUSD, book_btcusd_chan)
    c.WebSocket.AddSubscribe(bitfinex.CHAN_BOOK, bitfinex.LTCUSD, book_ltcusd_chan)
    c.WebSocket.AddSubscribe(bitfinex.CHAN_TRADE, bitfinex.BTCUSD, trades_chan)
    c.WebSocket.AddSubscribe(bitfinex.CHAN_TICKER, bitfinex.BTCUSD, ticker_chan)

    go listen(book_btcusd_chan, "BOOK BTCUSD:")
    go listen(book_ltcusd_chan, "BOOK LTCUSD:")
    go listen(trades_chan, "TRADES BTCUSD:")
    go listen(ticker_chan, "TICKER BTCUSD:")

    err = c.WebSocket.Subscribe()
    if err != nil {
        log.Fatal(err)
    }
}

func listen(in chan []float64, message string) {
    for {
        msg := <-in
        fmt.Println(message, msg)
    }
}
