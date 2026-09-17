package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestAncestors(t *testing.T) {
	assert.Strings(
		t,
		[]string{"doc/guide/alpha", "doc/guide", "doc"},
		pointer.Ancestors("doc/guide/alpha/reader/README.md"),
	)
	assert.Strings(t, nil, pointer.Ancestors("doc/README.md"))
	assert.Strings(t, nil, pointer.Ancestors(constant.ReadmeFile))
}
