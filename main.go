package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)
func main() {
	p := tea.newProgram(
		model{
			tracker: timer{},
		},
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong (err: %d)", err)
	}
}
