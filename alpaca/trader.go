package alpaca

import (
	"context"
	"log"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
)

const (
	LIVE_KEY_ENV_VAR = "ALPACA_LIVE_KEY"
	LIVE_SEC_ENV_VAR = "ALPACA_LIVE_SEC"
	PAPER_KEY_ENV_VAR = "ALPACA_PAPER_KEY"
	PAPER_SEC_ENV_VAR = "ALPACA_PAPER_SEC"
)

var Trader client

type client struct {
	toLIVE bool
	LIVE *alpaca.Client	
	PAPER *alpaca.Client	
}

func init() {
	live_key := os.Getenv(LIVE_KEY_ENV_VAR)
	live_sec := os.Getenv(LIVE_SEC_ENV_VAR)
	if live_key == "" || live_sec == "" {
		panic("[ALPACA-error] no key or secret setup for live trading")
	}

	paper_key := os.Getenv(PAPER_KEY_ENV_VAR)
	paper_sec := os.Getenv(PAPER_SEC_ENV_VAR)
	if paper_key == "" || paper_sec == "" {
		panic("[ALPACA-error] no key or secret setup for paper trading")
	}

	// Alternatively you can set your key and secret using the
	// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables

	live := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    live_key,
		APISecret: live_sec,
		BaseURL:   "https://api.alpaca.markets",
	})
	live.StreamTradeUpdatesInBackground(context.TODO(), func(tu alpaca.TradeUpdate) {
		log.Printf("[ALPACA-info] LIVE TRADE UPDATE: %+v\n", tu)
	})

	paper := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    paper_key,
		APISecret: paper_sec,
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	paper.StreamTradeUpdatesInBackground(context.TODO(), func(tu alpaca.TradeUpdate) {
		log.Printf("[ALPACA-info] PAPER TRADE UPDATE: %+v\n", tu)
	})

	Trader = client{
		LIVE: live,
		PAPER: paper,
	}
}

func (c *client) Buy(symbol string, qty uint64) (*alpaca.Order, error) {
	qtyDec := decimal.NewFromUint64(qty)
	return c.get().PlaceOrder(alpaca.PlaceOrderRequest{
		Symbol:      symbol,
		Qty:         &qtyDec,
		Side:        alpaca.Buy,
		Type:        alpaca.Market,
		TimeInForce: alpaca.Day})
}

func (c *client) get() *alpaca.Client {
	if c.toLIVE {
		return c.LIVE
	}
	return c.PAPER
}

