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
				sessionDone int
				onBreak bool
				onPause bool
			}{
				sessionDone: 0,
				onBreak: false,
				onPause: false,
			},
			pauseStart: 0,
			state: MAIN_MENU, 
			submenu: []submenu{ 
				{ 
					name: "Start",
					opts: []opt {
						{text: "Pause"},
						{text: "End Now"},
					},
				},
				{
					name: "Set work duration",
					opts: []opt {
						{text: "15"},
						{text: "30"},
						{text: "45"},
						{text: "60"},
					},
				},
				{
					name: "Set break duration",
					opts: []opt {
						{text: "5"},
						{text: "10"},
						{text: "15"},
						},
					},
				{
					name: "Set session",
					opts: []opt {
						{text: "2"},
						{text: "3"},
						{text: "4"},
						{text: "until tired"},
						},
					},
				{
					name: "Exit",
					opts: []opt{},
				},
			},
		},
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong (err: %d)", err)
	}
}
