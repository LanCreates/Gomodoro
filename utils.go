package main

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}

func tick() tea.Cmd {
	// Ticks 60 times per second
	return tea.Tick(
		time.Second / 60,
		func(time.Time) tea.Msg {
			return tickMsg{}
		},
	)
}

func getDigits(n int) int {
	digs := 0
	for; n > 0; n /= 10 { digs++ }
	return digs
}

func msToSecond(ms int64) int64 {
	return ms/1000
}

