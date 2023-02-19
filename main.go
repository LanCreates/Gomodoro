package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		model{ 
			config: struct{
				end int64
				session int
				workDuration int
				breakDuration int
			}{
				end: 0,
				session: 1,
				workDuration: 15,
				breakDuration: 5,
			}, 
			tracker: struct {
				onBreak bool
				onPause bool
			}{
				onBreak: false,
				onPause: false,
			},
			state: MAIN_MENU, 
			cursor: 0, 
			submenu: []submenu{ 
				{ 
					name: "Start",
					cursor: 0,
					opts: []opt {
						{text: "Pause"},
						{text: "End Now"},
					},
				},
				{
					name: "Set work duration",
					cursor: 0, 
					opts: []opt {
						{text: "15"},
						{text: "30"},
						{text: "45"},
						{text: "60"},
					},
				},
				{
					name: "Set break duration",
					cursor: 0, 
					opts: []opt {
						{text: "5"},
						{text: "10"},
						{text: "15"},
						},
					},
				{
					name: "Set session",
					cursor: 0, 
					opts: []opt {
						{text: "2"},
						{text: "3"},
						{text: "4"},
						{text: "5"},
						{text: "until tired"},
						},
					},
				{
					name: "Exit",
					cursor: 0, 
					opts: []opt{},
				},
			},
		},
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong (err: %d)", err)
	}
}
