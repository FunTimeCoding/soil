package unit_test

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
	// 1 space
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
		indenter.Parse("a\n b"),
	)
	// 2 space
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
		indenter.Parse("a\n  b"),
	)
	// 3 space
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
		indenter.Parse("a\n   b"),
	)
	// 4 space
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
		indenter.Parse("a\n    b"),
	)
	// 4 space with blank line
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
