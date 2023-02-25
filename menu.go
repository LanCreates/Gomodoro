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
	tracker struct {
		sessionDone int
		onBreak bool
		onPause bool
		pauseStart int64
	}
	state int
	cursor int
	submenu []submenu
}

func (m model) Init() tea.Cmd { 
	return tick()
}

// This handles state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 
	if m.state == EXIT {
		return m, tea.Quit
	}

	timeNow := time.Now().UnixMilli()
	if m.config.end - timeNow <= 0 {
		if m.state == BEGIN {
			if m.tracker.onBreak {
				notifyWork()
				m.config.end = time.Now().UnixMilli() + int64(m.config.workDuration * 1000 * 60)
				m.tracker.sessionDone++
			} else {
				m.config.end = int64(m.config.breakDuration * 1000 * 60)
				if m.tracker.sessionDone % 4 == 3 {
					notifyLongBreak()
					m.config.end *= 3
				} else {
					notifyBreak()
				}
				m.config.end += time.Now().UnixMilli()

			}

			m.tracker.onBreak = !(m.tracker.onBreak)
		}

		if m.tracker.sessionDone == m.config.session {
			m.state = MAIN_MENU
			m.tracker.sessionDone = 0
			m.tracker.onBreak = false
			m.tracker.onPause = false
		}
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
			m.config.breakDuration = 1
			m.config.workDuration = 1
			m.config.session = 4
		}
	case tickMsg:
		return m, tick()
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
	out = append(out, showControls())
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
		m.tracker.sessionDone = 0
		m.config.end = time.Now().UnixMilli() + int64(m.config.workDuration * 1000 * 60)
	}
}
