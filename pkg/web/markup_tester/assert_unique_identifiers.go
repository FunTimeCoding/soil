package markup_tester

import "github.com/PuerkitoBio/goquery"

func (o *Tester) AssertUniqueIdentifiers(body string) {
	o.t.Helper()
	seen := map[string]int{}
	document(body).Find("[id]").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			identifier, _ := s.Attr("id")
			seen[identifier]++
		},
	)

	for identifier, count := range seen {
		if count > 1 {
			o.t.Errorf("id %q appears %d times", identifier, count)
		}
	}
}
