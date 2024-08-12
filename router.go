package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rainbow-io-llc/tv-bot/handlers"
)

func register(r *gin.Engine) {
	r.GET("/ping", handlers.Ping())
	r.POST("/webhook", handlers.Webhook())
}