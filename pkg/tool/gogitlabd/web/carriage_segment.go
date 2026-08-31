package web

import "strings"

func carriageSegment(s string) string {
	s = strings.TrimSuffix(s, "\r")

	if index := strings.LastIndexByte(s, '\r'); index >= 0 {
		return s[index+1:]
	}

	return s
}
