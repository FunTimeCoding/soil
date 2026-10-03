package reference

import "fmt"

func Lines(findings []*Finding) []string {
	var result []string

	for _, f := range findings {
		result = append(result, fmt.Sprintf("  `%s` - %s", f.Span, f.Text))
	}

	return result
}
