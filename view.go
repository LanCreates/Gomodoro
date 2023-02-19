package main

import (
	"fmt"
	"strings"
)

func showConfig(m model) string {
	config := fmt.Sprintf(
		"Sessions: %d\nWork/Break time: %d/%d mins",
		m.config.session, m.config.workDuration, m.config.breakDuration,
	)
	return config
}
func viewMainMenu(m model) string {
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

func viewBegin(m model) string {
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
			out += fmt.Sprintf("%s %s", activeSelection(), v)
		} else {
			out += fmt.Sprintf("%s %s", nonactiveSelection(), v)
		}

		out += "    "
	}
	return out
}

func viewSetDurWork(m model) string {
	out := "How long do you want to work? (in minutes)\n"
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("%s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("%s %s", nonactiveSelection(), v.text)
		}
		out += "   "
	}
	return out
}

func viewSetDurBreak(m model) string {
	out := "How long do you want to break? (in minutes)\n"
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("%s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("%s %s", nonactiveSelection(), v.text)
		}
		out += "   "
	}
	return out
}

func viewSetSession(m model) string {
	out := "How many rounds do you want? (in minutes)\n"
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("%s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("%s %s", nonactiveSelection(), v.text)
		}
		out += "   "
	}
	return out
}
