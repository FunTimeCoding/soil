package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestChooseFix(t *testing.T) {
	t.Run(
		"SingleSegmentReturnsFirst",
		func(t *testing.T) {
			assert.String(
				t,
				"r",
				segment.ChooseFix("ref", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentPrefersMultiCharacter",
		func(t *testing.T) {
			assert.String(
				t,
				"reference",
				segment.ChooseFix("fooRef", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentFallsBackToFirst",
		func(t *testing.T) {
			assert.String(t, "t", segment.ChooseFix("fooTx", []string{"t"}))
		},
	)
	t.Run(
		"SingleOptionReturnsIt",
		func(t *testing.T) {
			assert.String(
				t,
				"arguments",
				segment.ChooseFix("args", []string{"arguments"}),
			)
		},
	)
}
