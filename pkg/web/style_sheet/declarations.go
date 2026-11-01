package style_sheet

import "strings"

func declarations(body string) []string {
	var result []string

	for _, d := range SplitTop(body, ';') {
		if d != "" {
			result = append(result, strings.Join(strings.Fields(d), " "))
		}
	}

	return result
}
