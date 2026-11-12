package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/goreach"
	"testing"
)

func TestTagsJoinsDiscoveredCustomTags(t *testing.T) {
	d := t.TempDir()
	testutil.WriteFile(
		t,
		d,
		"alfa/alfa.go",
		"//go:build local\n\npackage alfa\n",
	)
	testutil.WriteFile(t, d, "alfa/bravo.go", "//go:build ci\n\npackage alfa\n")
	testutil.WriteFile(
		t,
		d,
		"alfa/charlie.go",
		"//go:build linux\n\npackage alfa\n",
	)
	assert.String(t, "ci,local", goreach.Tags(d))
}
