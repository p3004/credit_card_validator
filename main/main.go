package main

import (
	"credit_card_validator/networking"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	networking.ValidateNumber(router)
	err := router.Run(networking.Host)
	if err != nil {
		return
	}
}
