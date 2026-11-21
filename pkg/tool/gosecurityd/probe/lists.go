package probe

import (
	"strconv"
	"strings"
	"time"
)

func lists(lines []string) *time.Time {
	if len(lines) == 0 {
		return nil
	}

	seconds, e := strconv.ParseFloat(strings.TrimSpace(lines[0]), 64)

	if e != nil || seconds <= 0 {
		return nil
	}

	result := time.Unix(int64(seconds), 0).UTC()

	return &result
}
