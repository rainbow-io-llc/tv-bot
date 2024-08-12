package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rainbow-io-llc/tv-bot/ngrok"
)

func main() {
	ctx := context.Background()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	
	// setup routers
	register(r)

	tunnel, err := ngrok.Tunnel.Get(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to create Ngrok Proxy: %s", err))
	}

	log.Printf("\n[GIN-info] spinning up server with tunnel...\n")
	if err := r.RunListener(tunnel); err != nil {
		panic(fmt.Sprintf("failed to spin up Gin server: %s", err))
	}
}
