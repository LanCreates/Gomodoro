package main

import (
	"strings"
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MAIN_MENU int = iota
	SET_DUR_WORK
	SET_DUR_BREAK
	SET_SESSION
	BEGIN
	EXIT
	N_MENU
)

type model struct {
	tracker struct{
		end int64
		session int
		workDuration int
		breakDuration int
		onBreak bool
	}
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
				if m.state != MAIN_MENU {
					m.navigateMenu(dir)
				} else {
					m.navigateSubmenu(dir)
				}
		}
	case tickMsg:
		
	}

	return m, nil
}

func (m model) View() string { 
	out := []string{ //╮╯─╰│╭╮
		"╭───────────────── Go-modoro ─────────────────╮",
	}

	switch m.state {
		case MAIN_MENU:
			out = append(out, viewMainMenu())
		case SET_DUR_WORK:
			out = append(out, viewSetDurWork())
		case SET_DUR_BREAK:
			out = append(out, viewSetDurBreak())
		case SET_SESSION:
			out = append(out, viewSetSession())
		case BEGIN:
			out = append(out, viewBegin())
	}
	out = append(out, "============================================")
	return strings.Join(out, "\n")
}

func (m *model) navigateMenu(dir string) {
	switch dir {
	case "up", "right": m.cursor++
	case "down", "left": m.cursor--
	}
	m.cursor = (m.cursor + N_MENU) % N_MENU
}

func (m *model) selectMenu() {
	m.state = m.cursor
	if m.state == BEGIN {
		m.tracker.end = time.Now().UnixMilli() + int64(m.tracker.workDuration * 1000)
	}
}
