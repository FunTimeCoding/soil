package system_page

import "strconv"

func ScoreValue(minimum float64) string {
	return strconv.FormatFloat(minimum, 'f', -1, 64)
}
