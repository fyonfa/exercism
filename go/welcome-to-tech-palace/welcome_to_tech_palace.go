package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsLine := starsPerLine(numStarsPerLine)

	return starsLine + "\n" + welcomeMsg + "\n" + starsLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	trimmedMsg := strings.Trim(oldMsg, "\n**\n\r\t ")
	return trimmedMsg
}

func starsPerLine(n int) string {
	symbol := "*"
	var starLines string
	for i := 0; i < n; i++ {
		starLines += fmt.Sprintf("%s", symbol)
	}
	return starLines
}
