package markup_tester

import "github.com/funtimecoding/soil/pkg/web/style_sheet"

func (o *Tester) AssertPicoWeight(body string) {
	o.t.Helper()
	d := document(body)
	reported := map[string]bool{}

	for _, r := range style_sheet.Parse(styleSheet(d)) {
		if ownVariablesOnly(r) {
			continue
		}

		for _, s := range r.Selectors {
			o.assertSelectorWeight(d, s, reported)
		}
	}
}
