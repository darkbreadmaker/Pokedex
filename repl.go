package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	var splitText []string = strings.Fields(lowerText)
	return splitText
}

