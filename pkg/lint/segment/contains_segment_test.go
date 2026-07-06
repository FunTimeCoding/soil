package segment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestContainsSegment(t *testing.T) {
	t.Run(
		"SingleLetterMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, ContainsSegment("r", "r"))
		},
	)
	t.Run(
		"SingleLetterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, ContainsSegment("fooBar", "r"))
		},
	)
	t.Run(
		"MultiCharacterPrefixMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, ContainsSegment("fooReference", "ref"))
		},
	)
	t.Run(
		"MultiCharacterExactMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, ContainsSegment("fooRef", "ref"))
		},
	)
	t.Run(
		"MultiCharacterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, ContainsSegment("fooBar", "ref"))
		},
	)
	t.Run(
		"MultiCharacterTooShort",
		func(t *testing.T) {
			assert.Boolean(t, false, ContainsSegment("fooRe", "ref"))
		},
	)
}
