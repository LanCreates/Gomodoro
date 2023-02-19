package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		model{ 
			tracker: timer{}, 
			state: MAIN_MENU, 
			cursor: 0, 
			submenu: []submenu{ 
				{cursor: 0, // SET WORK DURATION
				opts: []opt {
					{text: "15 min"},
					{text: "30 min"},
					{text: "45 min"},
					{text: "60 min"},
					},
				},
				{cursor: 0, // SET BREAK DURATION
				opts: []opt {
					{text: "5 min"},
					{text: "10 min"},
					{text: "15 min"},
					},
				},
				{cursor: 0, // SET SESSION
				opts: []opt {
					{text: "2"},
					{text: "3"},
					{text: "4"},
					{text: "5"},
					{text: "until tired"},
					},
				},
			},
		},
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong (err: %d)", err)
	}
}
