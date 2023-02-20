package main

import (
	"strings"
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	BEGIN int = iota
	SET_DUR_WORK 
	SET_DUR_BREAK
	SET_SESSION
	EXIT
	N_MENU
	MAIN_MENU
)

type model struct {
	config struct{
		end int64
		session int
		workDuration int
		breakDuration int
	}
	status struct {
		onBreak bool
		onPause bool
	}
	state int
	cursor int
	submenu []submenu
}

func (m model) Init() tea.Cmd { 
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 
	if m.config.end - time.Now().UnixMilli() < 0 {
		m.state = MAIN_MENU
	}

	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "enter":
			if m.state == MAIN_MENU {
				m.selectMenu()
			} else {
				m.selectSubmenu()
			}
		case "esc":
			if m.state != MAIN_MENU {
				m.state = MAIN_MENU
			}
		case "h", "j", "k", "l", "up", "down", "left", "right":
			dir := msg.(tea.KeyMsg).String()
			if m.state == MAIN_MENU {
				m.navigateMenu(dir)
			} else {
				m.navigateSubmenu(dir)
			}
		case "X":
			return m, tea.Quit
		}
	case tickMsg:
		
	}

	if m.state == EXIT {
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string { 
	out := []string{ //╮╯─╰│╭╮
		"╭───────────────── Go-modoro ─────────────────╮",
	}

	if m.state == MAIN_MENU {
		out = append(out, showConfig(m))
		out = append(out, "├─────────────────────────────────────────────┤")
	}

	switch m.state {
		case MAIN_MENU:
			out = append(out, viewMainMenu(m))
		case SET_DUR_WORK:
			out = append(out, viewSetDurWork(m))
		case SET_DUR_BREAK:
			out = append(out, viewSetDurBreak(m))
		case SET_SESSION:
			out = append(out, viewSetSession(m))
		case BEGIN:
			out = append(out, viewBegin(m))
	}
	out = append(out, "╰─────────────────────────────────────────────╯")
	return strings.Join(out, "\n")
}

func (m *model) navigateMenu(dir string) {
	switch dir {
	case "up", "right", "k", "l": m.cursor--
	case "down", "left", "h", "j": m.cursor++
	}
	m.cursor = (m.cursor + N_MENU) % N_MENU
}

func (m *model) selectMenu() {
	m.state = m.cursor
	if m.state == BEGIN {
		m.config.end = time.Now().UnixMilli() + int64(m.config.workDuration * 1000 * 60)
	}
}
