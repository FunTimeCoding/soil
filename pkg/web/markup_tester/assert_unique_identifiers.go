package markup_tester

import (
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func AssertUniqueIdentifiers(
	t *testing.T,
	body string,
) {
	t.Helper()
	seen := map[string]int{}
	document(body).Find("[id]").Each(
		func(_ int, s *goquery.Selection) {
			identifier, _ := s.Attr("id")
			seen[identifier]++
		},
	)

	for identifier, count := range seen {
		if count > 1 {
			t.Errorf("id %q appears %d times", identifier, count)
		}
	}
}
