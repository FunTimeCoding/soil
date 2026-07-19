package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.Any(
		t,
		[]string{"doc/ai/spec/naming.md"},
		pointer.Expand("doc/ai/spec/naming.md"),
	)
	assert.Any(
		t,
		[]string{"cmd/goalertlogd", "cmd/gomaintlogd"},
		pointer.Expand("cmd/{goalertlogd,gomaintlogd}"),
	)
	assert.Any(
		t,
		[]string{"a/c/f", "a/d/f", "b/c/f", "b/d/f"},
		pointer.Expand("{a,b}/{c,d}/f"),
	)
	assert.Any(t, []string{"doc/{only}/x"}, pointer.Expand("doc/{only}/x"))
	assert.Any(t, []string{"doc/{a,b"}, pointer.Expand("doc/{a,b"))
}
