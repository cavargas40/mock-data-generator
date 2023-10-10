package utils

import (
	"math/rand"
	"time"
)

// generates a random email address.
func GenerateRandomEmail() string {
	domains := []string{"gmail.com", "hotmail.com", "yahoo.com", "outlook.com"}
	rand.Seed(time.Now().UnixNano())
	randomDomain := domains[rand.Intn(len(domains))]
	randomPart := GenerateRandomString(10)
	return randomPart + "@" + randomDomain
}

// generates a random string of a given length.
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

// shuffles a string slice in place.
func ShuffleStringSlice(slice []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
