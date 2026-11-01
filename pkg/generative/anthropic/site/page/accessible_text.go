package page

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func accessibleText(s *goquery.Selection) string {
	c := s.Clone()
	c.Find(`[aria-hidden="true"]`).Remove()

	return strings.TrimSpace(c.Text())
}
