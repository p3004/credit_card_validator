package validator

import (
	"strconv"
)

func CardIssuerNameIdentifier(cardNumber string) string {
	mii, err := strconv.Atoi(string(cardNumber[0]))
	if err != nil {
		return ""
	}
	switch mii {
	case 4:
		return "Visa"
	case 5:
		return "Master Card"
	case 6:
		return "Rupay"
	default:
		return ""
	}
}

func IsValidCardNumber(creditCardNumber string) bool {
	var sum = 0
	length := len(creditCardNumber)
	even := true

	for i := length - 2; i >= 0; i-- {
		digit, err := strconv.Atoi(string(creditCardNumber[i]))
		if err != nil || digit < 0 || digit > 9 {
			return false
		}
		if even {
			digit *= 2
		}
		even = !even
		sum += digit/10 + digit%10
	}
	// checkSum is the last number of the credit card number, which we didn't use in the above loop
	checkSum, err := strconv.Atoi(string(creditCardNumber[length-1]))
	if err != nil {
		return false
	}
	return (sum+checkSum)%10 == 0
}
