package techpalace

import (
    "fmt"
    "strings"
)

func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

func CleanupMessage(oldMsg string) string {
    clean := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(clean)
}
