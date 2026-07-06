package indenter

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{Text: "a", Children: []*Node{}},
			},
		},
		Parse("a"),
	)
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{Text: "a", Children: []*Node{}},
				{Text: "b", Children: []*Node{}},
			},
		},
		Parse("a\nb"),
	)
	// 1 space
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{
					Text: "a",
					Children: []*Node{
						{Text: "b", Children: []*Node{}},
					},
				},
			},
		},
		Parse("a\n b"),
	)
	// 2 space
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{
					Text: "a",
					Children: []*Node{
						{Text: "b", Children: []*Node{}},
					},
				},
			},
		},
		Parse("a\n  b"),
	)
	// 3 space
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{
					Text: "a",
					Children: []*Node{
						{Text: "b", Children: []*Node{}},
					},
				},
			},
		},
		Parse("a\n   b"),
	)
	// 4 space
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{
					Text: "a",
					Children: []*Node{
						{Text: "b", Children: []*Node{}},
					},
				},
			},
		},
		Parse("a\n    b"),
	)
	// 4 space with blank line
	assert.Any(
		t,
		&Node{
			Children: []*Node{
				{
					Text: "a",
					Children: []*Node{
						{Text: "b", Children: []*Node{}},
					},
				},
			},
		},
		Parse("a\n\n    b"),
	)
}
