package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewSpan("doc/ai/spec/naming.md")},
		pointer.Expand(pointer.NewSpan("doc/ai/spec/naming.md")),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{
			pointer.NewSpan("cmd/goalertlogd"),
			pointer.NewSpan("cmd/gomaintlogd"),
		},
		pointer.Expand(pointer.NewSpan("cmd/{goalertlogd,gomaintlogd}")),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{
			pointer.NewLink("a/c/f"),
			pointer.NewLink("a/d/f"),
			pointer.NewLink("b/c/f"),
			pointer.NewLink("b/d/f"),
		},
		pointer.Expand(pointer.NewLink("{a,b}/{c,d}/f")),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewSpan("doc/{only}/x")},
		pointer.Expand(pointer.NewSpan("doc/{only}/x")),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewSpan("doc/{a,b")},
		pointer.Expand(pointer.NewSpan("doc/{a,b")),
	)
}
