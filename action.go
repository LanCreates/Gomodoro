package main

import (
	"strconv"
)

func (m *model) handleStart() {
}

func (m *model) handleSetDurWork(choice string) {
	m.config.workDuration, _ = strconv.Atoi(choice)
}

func (m *model) handleSetDurBreak(choice string) {
	m.config.breakDuration, _ = strconv.Atoi(choice)
}

func (m *model) handleSetSession(choice string) {

}
