package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/goreach"
	"testing"
)

func TestTagsJoinsDiscoveredCustomTags(t *testing.T) {
	d := t.TempDir()
	system_tester.New(t).Write(
		d,
		"alfa/alfa.go",
		"//go:build local\n\npackage alfa\n",
	)
	system_tester.New(t).Write(
		d,
		"alfa/bravo.go",
		"//go:build ci\n\npackage alfa\n",
	)
	system_tester.New(t).Write(
		d,
		"alfa/charlie.go",
		"//go:build linux\n\npackage alfa\n",
	)
	assert.String(t, "ci,local", goreach.Tags(d))
}
