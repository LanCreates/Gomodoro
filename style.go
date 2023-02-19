package main

import (
	"fmt"
	gloss "github.com/charmbracelet/lipgloss"
)

func decToHex(n int64) string {
	if(n < 0) { n += 256 }
	n %= 256
	
	digits := "0123456789ABCDEF"
	hex := ""
	for; n > 0; {
		hex = string(digits[n % 16]) + hex
		n /= 16
	}

	return hex
}

func makeColor(r, g, b int64) string {
	return "#" + decToHex(r) + decToHex(g) + decToHex(b)
}

var (
	ACTIVE_BLOCK = gloss.NewStyle().
		Background(gloss.Color("#FFFFFF"))
	NONACTIVE_BLOCK = gloss.NewStyle().
		Background(gloss.Color("#888888"))
)


func activeSelection() string {
	return fmt.Sprintf("%s",
		ACTIVE_BLOCK.Render("  "),
	)
}

func nonactiveSelection() string {
	return fmt.Sprintf("%s",
		NONACTIVE_BLOCK.Render("  "),
	)
}
