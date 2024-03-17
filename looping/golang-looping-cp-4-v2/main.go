package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	atIndex := strings.Index(email, "@")
	dotIndex := strings.Index(email, ".")

	provider := email[atIndex+1 : dotIndex]
	domain := email[dotIndex+1:]

	return fmt.Sprintf("Domain: %s dan TLD: %s", provider, domain)
}

func main() {
	// Test cases
	testCases := []string{"admin@yahoo.com", "ptmencaricintasejati@gmail.co.id"}

	for _, email := range testCases {
		fmt.Println(EmailInfo(email))
	}
}
