package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.Any(
		t,
		[]string{"doc/ai/spec/naming.md"},
		Expand("doc/ai/spec/naming.md"),
	)
	assert.Any(
		t,
		[]string{"cmd/goalertlogd", "cmd/gomaintlogd"},
		Expand("cmd/{goalertlogd,gomaintlogd}"),
	)
	assert.Any(
		t,
		[]string{"a/c/f", "a/d/f", "b/c/f", "b/d/f"},
		Expand("{a,b}/{c,d}/f"),
	)
	assert.Any(t, []string{"doc/{only}/x"}, Expand("doc/{only}/x"))
	assert.Any(t, []string{"doc/{a,b"}, Expand("doc/{a,b"))
}
