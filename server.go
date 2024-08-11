package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rainbow-io-llc/tv-bot/alpaca"
	"github.com/rainbow-io-llc/tv-bot/ngrok"
)

func main() {
	ctx := context.Background()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/ping", func(c *gin.Context) {
		order, err := alpaca.Trader.Buy("BTC/USD", 10)	
		if err != nil {
			log.Printf("\n[ALPACA-error] failed to place buy order: %s\n", err)
		} else {
			log.Printf("\n[ALPACA-info] place buy order successfully: %+v\n", *order)
		}
	})

	tunnel, err := ngrok.Tunnel.Get(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to create Ngrok Proxy: %s", err))
	}

	log.Printf("\n[GIN-info] spinning up server...\n")
	if err := r.RunListener(tunnel); err != nil {
		panic(fmt.Sprintf("failed to spin up server: %s", err))
	}
}
