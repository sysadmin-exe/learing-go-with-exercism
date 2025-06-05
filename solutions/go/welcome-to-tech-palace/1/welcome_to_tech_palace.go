package techpalace

import "strings"
import "fmt"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace, " + strings.ToUpper(customer) 
    return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
    messageWithBorder := border + "\n" + welcomeMsg + "\n" + border
    return messageWithBorder
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    starRemover := strings.NewReplacer("*", "")
    removeStars := starRemover.Replace(oldMsg)
	trimmedMessage := strings.TrimSpace(removeStars)
    fmt.Println(trimmedMessage)
    return trimmedMessage
}
