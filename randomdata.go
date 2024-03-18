package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomInteger generates a random integer between min and max (inclusive).
func GenerateRandomInteger(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GenerateRandomFloat generates a random float between min and max (inclusive).
func GenerateRandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

// GenerateRandomString generates a random string of given length.
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	// Generate a random integer between 1 and 100
	randomInt := GenerateRandomInteger(1, 100)
	fmt.Println("Random Integer:", randomInt)

	// Generate a random float between 0.0 and 1.0
	randomFloat := GenerateRandomFloat(0.0, 1.0)
	fmt.Println("Random Float:", randomFloat)

	// Generate a random string of length 10
	randomString := GenerateRandomString(10)
	fmt.Println("Random String:", randomString)
}
