package main

import "time"

type opt struct {
	text string
}

type submenu struct {
	name string
	cursor int
	opts []opt
}


func (m *model) navigateSubmenu(dir string) {
	switch dir {
	case "up", "right", "k", "l":
		m.submenu[m.cursor].cursor++
	case "down", "left", "h", "j":
		m.submenu[m.cursor].cursor--
	}
	nOpts := len(m.submenu[m.cursor].opts)
	m.submenu[m.cursor].cursor = (m.submenu[m.cursor].cursor + nOpts) % nOpts
}

func (m *model) selectSubmenu() {
	selected := m.submenu[m.cursor].cursor
	switch m.state {
	case BEGIN:
		switch selected {
		case 0: 
			m.tracker.onPause = !(m.tracker.onPause)
			if(m.tracker.onPause) {
				m.tracker.pauseStart = time.Now().UnixMilli()
			} else {
				if(m.config.end < time.Now().UnixMilli() - m.tracker.pauseStart) {
					m.config.end = time.Now().UnixMilli() + (m.config.end - m.tracker.pauseStart)
				} else {
					m.config.end += time.Now().UnixMilli() - m.tracker.pauseStart
				}
			}
		case 1: 
			m.tracker.onPause = false
			m.state = MAIN_MENU
		}
	case SET_DUR_WORK:
		m.config.workDuration = 25 + selected*10
		m.state = MAIN_MENU
	case SET_DUR_BREAK:
		m.config.breakDuration = (selected + 1)*5
		m.state = MAIN_MENU
	case SET_SESSION:
		if selected == 3 {
			m.config.session = 6969
		} else {
			m.config.session = (selected + 1)*8
		}
		m.state = MAIN_MENU
	}
}
