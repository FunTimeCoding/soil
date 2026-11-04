package basic

import "strings"

func paragraph(body []byte) string {
	s := string(body)
	start := strings.Index(s, "<p>")
	end := strings.Index(s, "</p>")

	if start < 0 || end <= start {
		return ""
	}

	return s[start+len("<p>") : end]
}
