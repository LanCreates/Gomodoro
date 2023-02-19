package main

import (
	"strconv"
)

func (m *model) handleStart() {

}

func (m *model) handleSetDurWork(choice string) {
	duration, _ := strconv.Atoi(choice)
	m.tracker.begin += int64(duration)
}

func (m *model) handleSetDurBreak(choice string) {
	duration, _ := strconv.Atoi(choice)
	m.tracker.begin += int64(duration)
}

func (m *model) handleSetSession(choice string) {

}
