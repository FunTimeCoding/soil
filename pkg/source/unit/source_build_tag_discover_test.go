package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func TestDiscoverLeavesOutUnixConstraint(t *testing.T) {
	d := t.TempDir()
	system_tester.New(t).Write(
		d,
		"alfa/alfa.go",
		"//go:build unix\n\npackage alfa\n",
	)
	system_tester.New(t).Write(
		d,
		"alfa/bravo.go",
		"//go:build browser\n\npackage alfa\n",
	)
	assert.Strings(t, []string{"browser"}, build_tag.Discover(d))
}

func TestDiscoverSkipsWhatGoToolSkips(t *testing.T) {
	d := t.TempDir()
	system_tester.New(t).Write(
		d,
		"tmp/alfa/alfa.go",
		"//go:build browser\n\npackage alfa\n",
	)
	system_tester.New(t).Write(
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
	system_tester.New(t).Write(
		d,
		"alfa/alfa.go",
		"//go:build browser\n\npackage alfa\n",
	)
	system_tester.New(t).Write(
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
