package main

import (
	"strconv"
)

func (m *model) handleStart() {
}

func (m *model) handleSetDurWork(choice string) {
	m.tracker.workDuration, _ = strconv.Atoi(choice)
}

func (m *model) handleSetDurBreak(choice string) {
	m.tracker.breakDuration, _ = strconv.Atoi(choice)
}

func (m *model) handleSetSession(choice string) {

}
