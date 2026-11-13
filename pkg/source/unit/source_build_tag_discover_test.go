package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
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

func TestDiscoverSkipsWhatGoToolSkips(t *testing.T) {
	d := t.TempDir()
	testutil.WriteFile(
		t,
		d,
		"tmp/alfa/alfa.go",
		"//go:build browser\n\npackage alfa\n",
	)
	testutil.WriteFile(
		t,
		d,
		"_probe/bravo.go",
		"//go:build local\n\npackage bravo\n",
	)
	assert.Strings(t, []string{"browser"}, build_tag.Discover(d))
	assert.Strings(
		t,
		[]string{"browser"},
		build_tag.Union(build_tag.TreeFiles(warm_tree.New(d))),
	)
}

func TestTreeFilesSkipTestdata(t *testing.T) {
	d := t.TempDir()
	testutil.WriteFile(
		t,
		d,
		"alfa/alfa.go",
		"//go:build browser\n\npackage alfa\n",
	)
	testutil.WriteFile(
		t,
		d,
		"alfa/testdata/bravo.go",
		"//go:build local\n\npackage bravo\n",
	)
	assert.Strings(
		t,
		[]string{"browser"},
		build_tag.Union(build_tag.TreeFiles(warm_tree.New(d))),
	)
}
