package markup_tester

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/soil/pkg/hypertext"
	"strings"
)

func document(body string) *goquery.Document {
	return hypertext.Document(strings.NewReader(body))
}
