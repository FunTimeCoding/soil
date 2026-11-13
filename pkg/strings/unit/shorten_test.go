package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/shorten"
	"testing"
)

func TestShortenPrefix(t *testing.T) {
	assert.String(
		t,
		"11111111",
		shorten.Prefix("11111111-2222-3333-4444-555555555555", 8),
	)
	assert.String(t, "exactly8", shorten.Prefix("exactly8", 8))
	assert.String(t, "short", shorten.Prefix("short", 8))
	assert.String(t, "", shorten.Prefix("", 8))
	assert.String(t, "", shorten.Prefix("anything", 0))
	assert.String(t, "äö", shorten.Prefix("äöü", 2))
}

func TestShortenEllipsisMarksCut(t *testing.T) {
	assert.String(t, "alfa…", shorten.Ellipsis("alfa bravo", 4))
	assert.String(t, "äö…", shorten.Ellipsis("äöü", 2))
}

func TestShortenEllipsisLeavesShortTextWhole(t *testing.T) {
	assert.String(t, "alfa", shorten.Ellipsis(constant.LowerAlfa, 4))
	assert.String(t, "", shorten.Ellipsis("", 4))
}
