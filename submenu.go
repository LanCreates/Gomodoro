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
	case "up", "right":
		m.submenu[m.cursor].cursor++
	case "down", "left":
		m.submenu[m.cursor].cursor--
	}
	m.submenu[m.cursor].cursor = (m.submenu[m.cursor].cursor + N_MENU) % N_MENU
}

