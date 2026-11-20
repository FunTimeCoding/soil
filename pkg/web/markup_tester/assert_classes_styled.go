package markup_tester

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"regexp"
	"strings"
)

func (o *Tester) AssertClassesStyled(body string) {
	o.t.Helper()
	d := document(body)
	styled := map[string]bool{}
	classSelector := regexp.MustCompile(`\.([a-zA-Z_][\w-]*)`)

	for _, s := range selectors(styleSheet(d)) {
		for _, m := range classSelector.FindAllStringSubmatch(s, -1) {
			styled[m[1]] = true
		}
	}

	for _, c := range constant.PicoClasses {
		styled[c] = true
	}

	reported := map[string]bool{}
	d.Find("[class]").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			value, _ := s.Attr("class")

			for _, c := range strings.Fields(value) {
				if !styled[c] && !reported[c] {
					reported[c] = true
					o.t.Errorf("class %q has no rule and is not Pico's", c)
				}
			}
		},
	)
}
