package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/scan"
	"testing"
)

func TestMatchFindsTermIgnoringCase(t *testing.T) {
	s := scan.New([]string{"Sprocket"})
	assert.Integer(t, 1, len(s.Match("the sprocket turns")))
	assert.Integer(t, 1, len(s.Match("THE SPROCKET TURNS")))
	assert.Integer(t, 0, len(s.Match("nothing here")))
}

func TestMatchRequiresWordBoundary(t *testing.T) {
	s := scan.New([]string{"widget"})
	assert.Integer(t, 0, len(s.Match("widgets are fine")))
	assert.Integer(t, 0, len(s.Match("a midwidget")))
	assert.Integer(t, 1, len(s.Match("a widget here")))
}

func TestMatchReachesIntoIdentifiers(t *testing.T) {
	s := scan.New([]string{"Sprocket"})
	assert.Integer(t, 1, len(s.Match("func SprocketHandler() {}")))
	assert.Integer(t, 1, len(s.Match("var innerSprocketCount int")))
}

func TestMatchReportsEveryTermOnTheLine(t *testing.T) {
	s := scan.New([]string{"widget", "Sprocket"})
	found := s.Match("widget and Sprocket together")
	assert.Integer(t, 2, len(found))
	assert.String(t, "widget", found[0])
	assert.String(t, "Sprocket", found[1])
}

func TestMatchDeduplicatesTermsIgnoringCase(t *testing.T) {
	s := scan.New([]string{"widget", "Widget", "WIDGET", ""})
	assert.Integer(t, 1, len(s.Match("one widget")))
}

func TestMatchHandlesNonAsciiTerms(t *testing.T) {
	s := scan.New([]string{"Grüße"})
	assert.Integer(t, 1, len(s.Match("die grüße kommen")))
	assert.Integer(t, 0, len(s.Match("die gruesse kommen")))
}

func TestMatchEscapesRegexMetacharacters(t *testing.T) {
	s := scan.New([]string{"a.b"})
	assert.Integer(t, 1, len(s.Match("the a.b value")))
	assert.Integer(t, 0, len(s.Match("the axb value")))
}
