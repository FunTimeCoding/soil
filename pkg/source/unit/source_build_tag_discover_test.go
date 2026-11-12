package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"testing"
)

func TestDiscoverLeavesOutUnixConstraint(t *testing.T) {
	d := t.TempDir()
	testutil.WriteFile(
		t,
		d,
		"alfa/alfa.go",
		"//go:build unix\n\npackage alfa\n",
	)
	testutil.WriteFile(
		t,
		d,
		"alfa/bravo.go",
		"//go:build browser\n\npackage alfa\n",
	)
	assert.Strings(t, []string{"browser"}, build_tag.Discover(d))
}
