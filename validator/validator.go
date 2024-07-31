package validator

import (
	"fmt"
	"strconv"
)

func IsValidCardNumber(creditCardNumber string) bool {
	var sum = 0
	length := len(creditCardNumber)
	even := true

	for i := length - 2; i >= 0; i-- {
		digit, err := strconv.Atoi(string(creditCardNumber[i]))
		if err != nil || digit < 0 || digit > 9 {
			return false
		}
		fmt.Println(digit)
		if even {
			digit *= 2
		}
		even = !even
		sum += digit/10 + digit%10
		fmt.Println(sum)
	}
	// checkSum is the last number of the credit card number, which we didn't use in the above loop
	checkSum, err := strconv.Atoi(string(creditCardNumber[length-1]))
	if err != nil {
		return false
	}
	return (sum+checkSum)%10 == 0
}
