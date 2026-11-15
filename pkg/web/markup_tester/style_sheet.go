package markup_tester

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func styleSheet(d *goquery.Document) string {
	var result strings.Builder
	d.Find("style").Each(
		func(_ int, s *goquery.Selection) {
			result.WriteString(s.Text())
		},
	)

	return result.String()
}
