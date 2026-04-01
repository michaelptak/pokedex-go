package main

import "strings"

func cleanInput(text string) []string {
	lowerStr := strings.ToLower(text)
	fields := strings.Fields(lowerStr)
	return fields
}
