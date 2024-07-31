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
		issuer := validator.CardIssuerNameIdentifier(request.CardNumber)
		if issuer == "" {
			context.JSON(http.StatusBadRequest, "We only accept card from Visa, MasterCard or Rupay")
			return
		}
		var response = Response{
			CardNumber: request.CardNumber,
			IsValid:    isValid,
			Issuer:     issuer,
		}
		context.JSON(http.StatusAccepted, response)
	})
}

func IssuerCheck(router *gin.Engine) {
	router.POST(issuerCheck, func(context *gin.Context) {
		var request Request
		err := context.BindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, "Unknown request body")
			return
		}
		issuer := validator.CardIssuerNameIdentifier(request.CardNumber)
		var issuerValidResponse = IssuerValid{
			IsValidIssuer: issuer != "",
			Issuer:        issuer,
		}
		context.JSON(http.StatusAccepted, issuerValidResponse)
	})
}
