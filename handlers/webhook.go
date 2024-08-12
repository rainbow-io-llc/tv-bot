package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	Symbol string `json:"symbol"`
	Action string `json:"action"`
	Time   string `json:"time"`
}

func Webhook() func(c *gin.Context) { 
	return func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("[Webhook-error] failed to bind body: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[Webhook-info] Received: %+v\n", payload)
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	}
}