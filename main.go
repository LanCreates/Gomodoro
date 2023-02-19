package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		model{ 
			tracker: struct{begin int64}{begin: 0}, 
			state: MAIN_MENU, 
			cursor: 0, 
			submenu: []submenu{ 
				{
					name: "Set work duration",
					cursor: 0, // SET WORK DURATION
					opts: []opt {
						{text: "15"},
						{text: "30"},
						{text: "45"},
						{text: "60"},
					},
				},
				{
				name: "Set break duration",
				cursor: 0, // SET BREAK DURATION
				opts: []opt {
					{text: "5"},
					{text: "10"},
					{text: "15"},
					},
				},
				{
				name: "Set session",
				cursor: 0, // SET SESSION
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
