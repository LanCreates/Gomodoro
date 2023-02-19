package main

type opt struct {
	text string
}

type submenu struct {
	name string
	cursor int
	opts []opt
}

type tickMsg struct{}

func (m *model) navigateSubmenu(dir string) {
	switch dir {
	case "up", "right", "k", "l":
		m.submenu[m.cursor].cursor++
	case "down", "left", "h", "j":
		m.submenu[m.cursor].cursor--
	}
	m.submenu[m.cursor].cursor = (m.submenu[m.cursor].cursor + N_MENU) % N_MENU
}

func (m model) selectSubmenu() {
	switch m.state {
		case MAIN_MENU:
		case SET_DUR_WORK:
		case SET_DUR_BREAK:
		case SET_SESSION:
		case BEGIN:
	}
}
