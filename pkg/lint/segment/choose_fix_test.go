package segment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestChooseFix(t *testing.T) {
	t.Run(
		"SingleSegmentReturnsFirst",
		func(t *testing.T) {
			assert.String(
				t,
				"r",
				ChooseFix("ref", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentPrefersMultiCharacter",
		func(t *testing.T) {
			assert.String(
				t,
				"reference",
				ChooseFix("fooRef", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentFallsBackToFirst",
		func(t *testing.T) {
			assert.String(t, "t", ChooseFix("fooTx", []string{"t"}))
		},
	)
	t.Run(
		"SingleOptionReturnsIt",
		func(t *testing.T) {
			assert.String(
				t,
				"arguments",
				ChooseFix("args", []string{"arguments"}),
			)
		},
	)
}
