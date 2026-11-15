package markup_tester

import (
	"github.com/funtimecoding/soil/pkg/web/style_sheet"
	"testing"
)

func AssertPicoWeight(
	t *testing.T,
	body string,
) {
	t.Helper()
	d := document(body)
	reported := map[string]bool{}

	for _, r := range style_sheet.Parse(styleSheet(d)) {
		if ownVariablesOnly(r) {
			continue
		}

		for _, s := range r.Selectors {
			assertSelectorWeight(t, d, s, reported)
		}
	}
}
