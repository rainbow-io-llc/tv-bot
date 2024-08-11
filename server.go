package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	ctx := context.Background()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	tunnel, err := getTunnel(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to create Ngrok Proxy: %s", err))
	}
	log.Printf("\n[NGROK-info] ingress URL: %s\n", tunnel.Addr())

	log.Printf("\n[GIN-info] spinning up server...\n")
	if err := r.RunListener(tunnel); err != nil {
		panic(fmt.Sprintf("failed to spin up server: %s", err))
	}
}

func getTunnel(ctx context.Context) (net.Listener, error) {
	username := os.Getenv("NGROK_USRNAME")
	password := os.Getenv("NGROK_PASSWD")
	
	log.Printf("\n[NGROK-info] spinning up tunnel...\n")
	return ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithBasicAuth(username, password),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
}
