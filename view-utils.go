package main

import "strconv"

func updateSegmentDisplay(timeLeft int) {
	timeStr := strconv.Itoa(timeLeft / 60)
	if len(timeStr) == 1 {
		timeStr = "0" + timeStr
	}

	timeStr += strconv.Itoa(timeLeft % 60)
	if len(timeStr) == 3 {
		timeStr = timeStr[:2] + "0" + string(timeStr[2])
	}

	for ix, v := range(timeStr) {
		for jx := 0; jx < 7; jx++ {
			segmentDisplay[ix][jx] = false
		}

		switch v {
		case '1':
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][5] = true
		case '2':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][4] = true
			segmentDisplay[ix][6] = true
		case '3':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		case '4':
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][5] = true
		case '5':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		case '6':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][4] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		case '7':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][5] = true
		case '8':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][4] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		case '9':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][3] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		case '0':
			segmentDisplay[ix][0] = true
			segmentDisplay[ix][1] = true
			segmentDisplay[ix][2] = true
			segmentDisplay[ix][4] = true
			segmentDisplay[ix][5] = true
			segmentDisplay[ix][6] = true
		}
	}
}
