package main

import (
	"fmt"
	"strings"
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
	BLOCK_ACTIVE = gloss.NewStyle().
		Background(gloss.Color("#FAA834"))
	BLOCK_BRIGHT = gloss.NewStyle().
		Background(gloss.Color("#EEEEEE"))
	BLOCK_DIM = gloss.NewStyle().
		Background(gloss.Color("#666666"))
	TXT_DIM = gloss.NewStyle().
		Foreground(gloss.Color("#888888"))
	TXT_ONSELECT = gloss.NewStyle().
		Foreground(gloss.Color("#FAA834"))
)


func activeBlock(l int) string {
	return fmt.Sprintf("%s",
		BLOCK_ACTIVE.Render(strings.Repeat(" ", l)),
	)
}

func dimBlock(l int) string {
	return fmt.Sprintf("%s",
		BLOCK_DIM.Render(strings.Repeat(" ", l)),
	)
}

func dimText(s string) string {
	return fmt.Sprintf("%s",
		TXT_DIM.Render(s),
	)
}
