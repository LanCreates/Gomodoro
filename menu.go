package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MAIN_MENU int = iota
	SET_DUR_BREAK
	SET_DUR_WORK
	SET_SESSION
	TIMER
	N_MENU
)

type model struct {
	tracker struct{begin int64}
	state int
	cursor int
	submenu []submenu
}

func (m model) Init() tea.Cmd { 
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
			case "esc":
				return m, tea.Quit
			case "up", "down", "left", "right":
		}
	case tickMsg:
		
	}

	return m, nil
}

func (m model) View() string { 
	out := ""
	return out
}

func (m *model) navigateMenu(dir string) {
	switch dir {
	case "up", "right": m.cursor++
	case "down", "left": m.cursor--
	}
	m.cursor = (m.cursor + N_MENU) % N_MENU
}

