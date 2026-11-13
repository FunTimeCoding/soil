package search_page

import "strings"

func (p *Page) terms() []string {
	var result []string

	for _, word := range strings.Fields(p.query) {
		if strings.HasPrefix(word, "-") {
			continue
		}

		trimmed := strings.Trim(word, `"'`)

		if len(trimmed) > 1 {
			result = append(result, trimmed)
		}
	}

	return result
}
