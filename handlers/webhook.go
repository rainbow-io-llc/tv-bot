package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/rainbow-io-llc/tv-bot/broker/alpaca"
)

const TOKEN_ENV_VAR = "WEBHOOK_TOKEN"

var token string

type Payload struct {
	Symbol string `json:"symbol"`
	Action string `json:"action"`
	Price  string `json:"price"`
	Time   string `json:"time"`
	Token  string `json:"token"`
}

func init() {
	token = os.Getenv(TOKEN_ENV_VAR)
	if token == "" {
		panic("[WEBHOOK-error] no token setup for tunnel connect")
	}
}

func Webhook() func(c *gin.Context) { 
	return func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("[WEBHOOK-error] failed to bind body: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("[WEBHOOK-info] Received: %+v\n", payload)

		if payload.Token != token {
			log.Printf("[WEBHOOK-error] invalid token in payload: %s\n", payload.Token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("invalid token")})
			return	
		}

		if payload.Action == "buy" {
			// TODO: use 100 qty as default here, change to percentage of portfolio
			if _, err := alpaca.Trader.Buy(payload.Symbol, 100); err != nil {
				log.Printf("[WEBHOOK-error] failed to execute buy: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			if err := alpaca.Trader.Sell(payload.Symbol); err != nil {
				log.Printf("[WEBHOOK-error] failed to execute sell: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	}
}