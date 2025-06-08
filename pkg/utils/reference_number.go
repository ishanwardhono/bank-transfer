package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateReferenceNumber() string {
	now := time.Now()
	dateStr := now.Format("20060102")
	randomStr := generateRandomString(8)
	return fmt.Sprintf("TRF-%s-%s", dateStr, randomStr)
}

func generateRandomString(length int) string {
	// Create a new random source with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define characters to use in random string
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var sb strings.Builder
	sb.Grow(length)

	for range length {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}

	return sb.String()
}
