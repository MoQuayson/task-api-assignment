package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	const tokenLength = 32
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	token := make([]byte, tokenLength)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func GenerateMobileNumber() string {
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < 10; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}

	return result
}

func GenerateRandomAmount() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}

func GenerateTimeSleepDuration() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(30-15+1) + 15)

}

func RandomizePaymentStatus() PaymentStatus {
	val := GenerateRandomAmount()

	if val%2 == 0 {
		return PaymentStatus_Successful
	}

	return PaymentStatus_Failed
}
