package markup_tester

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"slices"
	"strings"
)

func (o *Tester) assertSelectorWeight(
	d *goquery.Document,
	s string,
	reported map[string]bool,
) {
	o.t.Helper()
	weight := Specificity(s)

	if strings.Contains(s, "::") || weight != [3]int{} {
		return
	}

	for _, a := range Alternatives(s) {
		m, e := cascadia.Compile(a)

		if e != nil {
			o.t.Errorf("rule %q: alternative %q does not compile: %v", s, a, e)

			continue
		}

		d.FindMatcher(m).Each(
			func(
				_ int,
				n *goquery.Selection,
			) {
				tag := goquery.NodeName(n)
				key := join.Empty(s, "|", tag)

				if slices.Contains(
					constant.PicoElements,
					tag,
				) && !reported[key] {
					reported[key] = true
					o.t.Errorf(
						"rule %q weighs %v on <%s>, which Pico styles",
						s,
						weight,
						tag,
					)
				}
			},
		)
	}
}
