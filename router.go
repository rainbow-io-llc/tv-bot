package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rainbow-io-llc/tv-bot/alpaca"
)

func register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pong",
		})
	  })

	r.POST("/webhook", func(c *gin.Context) {
		assets, err := alpaca.Trader.Check("AAPL")	
		if err != nil {
			log.Printf("\n[ALPACA-error] failed to check orders: %s\n", err)
		} 
		if assets != nil {
			log.Printf("\n[ALPACA-info] orders: %+v\n", assets)	
		}
	})
}