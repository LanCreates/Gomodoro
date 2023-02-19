package main

import (
	"fmt"
	"strings"
)

func (m model) viewMainMenu() string {
	var out []string
	for k, v := range(m.submenu) {
		if k == m.cursor {
			out = append(out, fmt.Sprintf(" >  %s", v.name))
		} else {
			out = append(out, fmt.Sprintf("   %s", v.name))
		}
	}
	return strings.Join(out, "\n")
}

func (m model) viewBegin() string {
	out := ""

	for k, v := range([]string{"", "That's a Wrap..."}) {
		if k == 0 {
			if m.status.onPause {
				v = "Let's Continue!  "
			} else {
				v = "Hold on a second!"
			}
		}

		if(k == m.submenu[m.cursor].cursor) {
			out += fmt.Sprintf(" > %s", v)
		} else {
			out += fmt.Sprintf("   %s", v)
		}

		out += "    "
	}
	return out
}

func viewSetDurWork() string {
	out := "Set Work"
	return out
}

func viewSetDurBreak() string {
	out := "Set Break"
	return out
}

func viewSetSession() string {
	out := "Set Session"
	return out
}
