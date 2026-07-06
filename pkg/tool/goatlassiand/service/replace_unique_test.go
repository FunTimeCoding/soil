package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReplaceUnique(t *testing.T) {
	cases := []struct {
		name    string
		content string
		old     string
		new     string
		expect  string
		fail    string
	}{
		{
			"single match",
			"Hello world",
			"world",
			"there",
			"Hello there",
			"",
		},
		{
			"not found",
			"Hello world",
			"missing",
			"replacement",
			"",
			"old_text not found in page",
		},
		{
			"multiple matches",
			"foo bar foo baz foo",
			"foo",
			"qux",
			"",
			"old_text found 3 times, must be unique",
		},
		{
			"two matches",
			"the cat and the dog",
			"the",
			"a",
			"",
			"old_text found 2 times, must be unique",
		},
		{
			"match at start",
			"## Heading\n\nBody text",
			"## Heading",
			"## New Heading",
			"## New Heading\n\nBody text",
			"",
		},
		{
			"match at end",
			"First paragraph\n\nLast line",
			"Last line",
			"Final line",
			"First paragraph\n\nFinal line",
			"",
		},
		{
			"multiline old text",
			"# Title\n\n- Alfa\n- Bravo\n\nEnd",
			"- Alfa\n- Bravo",
			"- Charlie\n- Delta\n- Echo",
			"# Title\n\n- Charlie\n- Delta\n- Echo\n\nEnd",
			"",
		},
		{
			"replace with empty",
			"Keep this. Remove this. Keep that.",
			" Remove this.",
			"",
			"Keep this. Keep that.",
			"",
		},
		{
			"replace empty section with content",
			"Before\n\nAfter",
			"Before\n\nAfter",
			"Before\n\nNew middle section\n\nAfter",
			"Before\n\nNew middle section\n\nAfter",
			"",
		},
		{
			"unique by context",
			"the cat sat on the mat",
			"the cat",
			"a dog",
			"a dog sat on the mat",
			"",
		},
	}

	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				result, e := replaceUnique(c.content, c.old, c.new)

				if c.fail != "" {
					if e == nil {
						t.Fatalf("expected error containing %q", c.fail)
					}

					assert.StringContains(t, c.fail, e.Error())

					return
				}

				if e != nil {
					t.Fatalf("unexpected error: %v", e)
				}

				assert.String(t, c.expect, result)
			},
		)
	}
}
