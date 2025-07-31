package core

import (
	"fmt"
	"strings"
)

func UserInput(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func UserInputToBool(s string) bool {
	lower := strings.ToLower(strings.TrimSpace(s))

	switch lower {
	case "y", "yes":
		return true
	default:
		return false
	}
}
