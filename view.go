package main

import (
	"fmt"
	"time"
	"strings"
)

var segmentDisplay = [4][7]bool{}

func showControls() string {
	out := []string{dimText("                   Controls                   ")}

	for _, v := range([]string{"^/k - up               v/j - down", "</h - left             >/l - right"}) {
		out = append(out, dimText(v))
	}

	return strings.Join(out, "\n      ")
}

func showConfig(m model) string {
	nSessions := m.config.session
	workDur, breakDur := m.config.workDuration, m.config.breakDuration

	config := fmt.Sprintf(
		"│   Sessions: %d%s│\n│   Work/Break time: %d/%d mins%s│",
		nSessions, 
		strings.Repeat(" ", 45 - (12 + getDigits(nSessions))),
		workDur, 
		breakDur, 
		strings.Repeat(" ", 45 - (18 + getDigits(workDur) + getDigits(breakDur) + 6)),
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
					temp = append(temp, fmt.Sprintf(" %s ", strings.Repeat("8", 6)))
				} else {
					temp = append(temp, fmt.Sprintf(" %s ", strings.Repeat(dimText("."), 6)))
				}
			}
			out += "│  " + strings.Join(temp, "   ") + "  │\n"
		} else {
			for digit := 0; digit < 4; digit++ {
				left, right := dimText("."), dimText(".")
				if(segmentDisplay[digit][ix] == true) { left = "8" }
				if(segmentDisplay[digit][ix + 1] == true) { right = "8" }
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
					dimText(v.name), strings.Repeat(" ", 45 - (3 + len(v.name))),
				),
			)
		}
	}
	return strings.Join(out, "\n")
}

func viewBegin(m model) string {
	out := showTimer(m)
	out += "│        Minutes               Seconds        │\n"
	out += "├─────────────────────────────────────────────┤\n"

	if m.tracker.onBreak {
		if m.tracker.sessionDone % 4 == 3 {
			out += "│ Longer Break...   "
		} else {
			out += "│ Break...          "
		}
	} else {
		out += "│ Working...        "
	}

	if m.tracker.onPause {
		out += "                    Pause │\n"
	} else {
		out += "                  Running │\n"
	}

	out += fmt.Sprintf("│ Session %d/%d%s│\n", 
		m.tracker.sessionDone, m.config.session,
		strings.Repeat(" ", 
			45 - (9 + getDigits(m.tracker.sessionDone) + getDigits(m.config.session)),
		),
	)

	out += "├─────────────────────────────────────────────┤\n"
	for k, v := range([]string{"", "That's a Wrap..."}) {
		if k == 0 {
			if m.tracker.onPause {
				v = "Let's Continue!  "
			} else {
				v = "Hold on a second!"
			}
		}

		if(k == m.submenu[m.cursor].cursor) {
			out += fmt.Sprintf("│ > %s%s│", 
				v, strings.Repeat(" ", 45 - (3 + len(v))),
			)
		} else {
			out += fmt.Sprintf("│  %s%s│", 
				dimText(v), strings.Repeat(" ", 45 - (2 + len(v))),
			)
		}
		
		if k == 0 {out += "\n"}
	}
	return out
}

func viewSetDurWork(m model) string {
	out := "│  How long do you want to work? (in minute)  │\n│    "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeBlock(2), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", dimBlock(2), dimText(v.text))
		}
	}
	return out + "         │"
}

func viewSetDurBreak(m model) string {
	out := "│  How long do you want to rest? (in minute)  │\n│          "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeBlock(2), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", dimBlock(2), dimText(v.text))
		}
	}
	return out + "            │"
}

func viewSetSession(m model) string {
	out := "│        How many rounds do you want?         │\n│  "
	for k, v := range(m.submenu[m.cursor].opts) {
		if k == m.submenu[m.cursor].cursor {
			out += fmt.Sprintf("   %s %s", activeBlock(2), v.text)
		} else {
			out += fmt.Sprintf("   %s %s", dimBlock(2), dimText(v.text))
		}
	}
	return out + "   │"
}
