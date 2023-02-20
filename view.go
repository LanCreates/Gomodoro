package main

import (
	"fmt"
	"strings"
)

func getDigits(n int) int {
	digs := 0
	for; n > 0; n /= 10 { digs++ }
	return digs
}

func showConfig(m model) string {
	nSessions := m.config.session
	workDur, breakDur := m.config.workDuration, m.config.breakDuration

	// Spaces determined, do not alter!
	config := fmt.Sprintf(
		"│   Sessions: %d%s│\n│   Work/Break time: %d/%d mins%s│",
		nSessions, 
		strings.Repeat(" ", 45 - (13 + getDigits(nSessions))),
		workDur, 
		breakDur, 
		strings.Repeat(" ", 45 - (20 + getDigits(workDur) + getDigits(breakDur) + 6)),
	)
	return config
}

func viewMainMenu(m model) string {
	var out []string
	for k, v := range(m.submenu) {
		if k == m.cursor {
			out = append(out, 
				fmt.Sprintf( "│ >  %s%s│", 
					v.name, strings.Repeat(" ", 45 - (4 + len(v.name))),
				),
			)
		} else {
			out = append(out, 
				fmt.Sprintf( "│   %s%s│", 
					v.name, strings.Repeat(" ", 45 - (3 + len(v.name))),
				),
			)
		}
	}
	return strings.Join(out, "\n")
}

func viewBegin(m model) string {
	out := "│"

	for k, v := range([]string{"", "That's a Wrap..."}) {
		if k == 0 {
			if m.status.onPause {
				v = "Let's Continue!  "
			} else {
				v = "Hold on a second!"
			}
		}

		if(k == m.submenu[m.cursor].cursor) {
			out += fmt.Sprintf("  %s %s", activeSelection(), v)
		} else {
			out += fmt.Sprintf("  %s %s", nonactiveSelection(), v)
		}
	}
	return out + "  │"
}

func viewSetDurWork(m model) string {
	out := "│  How long do you want to work? (in minute)  │\n│    "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", nonactiveSelection(), v.text)
		}
	}
	return out + "         │"
}

func viewSetDurBreak(m model) string {
	out := "│  How long do you want to rest? (in minute)  │\n│          "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", nonactiveSelection(), v.text)
		}
	}
	return out + "            │"
}

func viewSetSession(m model) string {
	out := "│        How many rounds do you want?         │\n│  "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeSelection(), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", nonactiveSelection(), v.text)
		}
	}
	return out + "     │"
}

func showControls() string {
	out := ""
	return out
}
