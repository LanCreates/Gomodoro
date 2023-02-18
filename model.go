package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type timer struct {
	begin int64
}

type model struct {
	tracker timer
}

func (m model) Init() tea.Cmd { 
	return nil
}

func (m model) Update() (tea.Model, tea.Cmd) { 
	return m, nil
}

func (m model) View() string { 
	out := ""
	return out
}
