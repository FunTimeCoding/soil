package style_sheet

import "strings"

func blankComments(css string) string {
	b := []byte(css)

	for {
		start := strings.Index(string(b), "/*")

		if start < 0 {
			return string(b)
		}

		end := strings.Index(string(b[start:]), "*/")

		if end < 0 {
			end = len(b) - start
		} else {
			end += 2
		}

		for i := start; i < start+end; i++ {
			if b[i] != '\n' {
				b[i] = ' '
			}
		}
	}
}
