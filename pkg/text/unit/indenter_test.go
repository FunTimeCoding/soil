package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/indenter"
	"testing"
)

func TestIndenterParse(t *testing.T) {
	assert.Any(
		t,
		&indenter.Node{
			Children: []*indenter.Node{
				{Text: "a", Children: []*indenter.Node{}},
			},
		},
		indenter.Parse("a"),
	)
	assert.Any(
		t,
		&indenter.Node{
			Children: []*indenter.Node{
				{Text: "a", Children: []*indenter.Node{}},
				{Text: "b", Children: []*indenter.Node{}},
			},
		},
		indenter.Parse("a\nb"),
	)
}

func TestIndenterParseNestsAtAnyIndentWidth(t *testing.T) {
	for _, input := range []string{
		"a\n b",
		"a\n  b",
		"a\n   b",
		"a\n    b",
	} {
		assert.Any(
			t,
			&indenter.Node{
				Children: []*indenter.Node{
					{
						Text: "a",
						Children: []*indenter.Node{
							{Text: "b", Children: []*indenter.Node{}},
						},
					},
				},
			},
			indenter.Parse(input),
		)
	}
}

func TestIndenterParseNestsAcrossBlankLine(t *testing.T) {
	assert.Any(
		t,
		&indenter.Node{
			Children: []*indenter.Node{
				{
					Text: "a",
					Children: []*indenter.Node{
						{Text: "b", Children: []*indenter.Node{}},
					},
				},
			},
		},
		indenter.Parse("a\n\n    b"),
	)
}
