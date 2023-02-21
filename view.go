package main

import (
	"fmt"
	"time"
	"strings"
)

var segmentDisplay = [4][7]bool{}

func showControls() string {
	out := []string{"                   Controls                   "}

	for _, v := range([]string{"^/k - up               v/j - down", "</h - left             >/l - right"}) {
		out = append(out, v)
	}

	return strings.Join(out, "\n      ")
}

func showConfig(m model) string {
	nSessions := m.config.session
	workDur, breakDur := m.config.workDuration, m.config.breakDuration

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

func showTimer(m model) string {
	timeLeft := int(msToSecond(m.config.end - time.Now().UnixMilli()))
	if !(m.tracker.onPause) {
		updateSegmentDisplay(timeLeft)
	}

	var out string
	for ix := 0; ix < 7; ix++ {
		var temp []string = []string{}
		if ix % 3 == 0 {
			for digit := 0; digit < 4; digit++ {
				if(segmentDisplay[digit][ix] == true) {
					temp = append(temp, fmt.Sprintf(" %s ", strings.Repeat("0", 6)))
				} else {
					temp = append(temp, fmt.Sprintf(" %s ", strings.Repeat(".", 6)))
				}
			}
			out += "│  " + strings.Join(temp, "   ") + "  │\n"
		} else {
			for digit := 0; digit < 4; digit++ {
				left, right := ".", "."
				if(segmentDisplay[digit][ix] == true) { left = "0" }
				if(segmentDisplay[digit][ix + 1] == true) { right = "0" }
				temp = append(temp, fmt.Sprintf("%s%s%s", left, strings.Repeat(" ", 6), right))
			}

			for row := 0; row < 3; row++ {
				out += "│  " + strings.Join(temp, "   ") + "  │\n"
			}

			ix++
		}
	}
	return out
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
	out := showTimer(m)
	out += "│        Minutes               Seconds        │\n"
	if m.tracker.onPause {
		out += "│                    Pause                    │\n│"
	} else {
		out += "│                  Running..                  │\n│"
	}
	for k, v := range([]string{"", "That's a Wrap..."}) {
		if k == 0 {
			if m.tracker.onPause {
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
	return out + "   │"
}
