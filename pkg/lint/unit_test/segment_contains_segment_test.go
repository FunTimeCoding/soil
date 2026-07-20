package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestContainsSegment(t *testing.T) {
	t.Run(
		"SingleLetterMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, segment.ContainsSegment("r", "r"))
		},
	)
	t.Run(
		"SingleLetterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, segment.ContainsSegment("fooBar", "r"))
		},
	)
	t.Run(
		"MultiCharacterPrefixMatch",
		func(t *testing.T) {
			assert.Boolean(
				t,
				true,
				segment.ContainsSegment("fooReference", "ref"),
			)
		},
	)
	t.Run(
		"MultiCharacterExactMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, segment.ContainsSegment("fooRef", "ref"))
		},
	)
	t.Run(
		"MultiCharacterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, segment.ContainsSegment("fooBar", "ref"))
		},
	)
	t.Run(
		"MultiCharacterTooShort",
		func(t *testing.T) {
			assert.Boolean(t, false, segment.ContainsSegment("fooRe", "ref"))
		},
	)
}
