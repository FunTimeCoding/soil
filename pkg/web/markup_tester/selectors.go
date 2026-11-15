package markup_tester

import "github.com/funtimecoding/soil/pkg/web/style_sheet"

func selectors(css string) []string {
	var result []string

	for _, r := range style_sheet.Parse(css) {
		result = append(result, r.Selectors...)
	}

	return result
}
