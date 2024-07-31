package main

import (
	"credit_card_validator/validator"
	"fmt"
	"strconv"
)

var creditCardNumber string

func main() {
	fmt.Println("Enter your credit card number: ")
	_, err := fmt.Scanf("%s", &creditCardNumber)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	isValid := validator.IsValidCardNumber(creditCardNumber)
	fmt.Println("Credit card number " + creditCardNumber + " is " + strconv.FormatBool(isValid))
}
