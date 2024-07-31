package networking

import (
	"credit_card_validator/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateNumber(router *gin.Engine) {
	router.POST(validate, func(context *gin.Context) {
		var request Request
		if err := context.ShouldBindBodyWithJSON(&request); err != nil {
			return
		}
		isValid := validator.IsValidCardNumber(request.CardNumber)
		var response = Response{
			CardNumber: request.CardNumber,
			IsValid:    isValid,
		}
		context.JSON(http.StatusAccepted, response)
	})
}
