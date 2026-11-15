package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/unit/go_mod_tester"
	"golang.org/x/mod/modfile"
	"testing"
)

func TestSyncReplacesAddsManagedBlock(t *testing.T) {
	assert.String(
		t,
		"module alfa\n\ngo 1.27\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.0.0\n\tcharlie.test/c => charlie.test/c v2.0.0\n)\n",
		go_mod_tester.Sync(
			t,
			"module alfa\n\ngo 1.27\n",
			go_mod_tester.Replace("bravo.test/b", "v1.0.0"),
			go_mod_tester.Replace("charlie.test/c", "v2.0.0"),
		),
	)
}

func TestSyncReplacesKeepsHandReplaces(t *testing.T) {
	assert.String(
		t,
		"module alfa\n\ngo 1.27\n\nreplace delta.test/d => ../delta\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.1.0\n)\n",
		go_mod_tester.Sync(
			t,
			"module alfa\n\ngo 1.27\n\nreplace delta.test/d => ../delta\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.0.0\n\tcharlie.test/c => charlie.test/c v2.0.0\n)\n",
			go_mod_tester.Replace("bravo.test/b", "v1.1.0"),
		),
	)
}

func TestSyncReplacesKeepsTrailingDirectives(t *testing.T) {
	assert.String(
		t,
		"module alfa\n\ngo 1.27\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.1.0\n)\n\nretract v0.1.0\n",
		go_mod_tester.Sync(
			t,
			"module alfa\n\ngo 1.27\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.0.0\n)\n\nretract v0.1.0\n",
			go_mod_tester.Replace("bravo.test/b", "v1.1.0"),
		),
	)
}

func TestSyncReplacesStable(t *testing.T) {
	first := go_mod_tester.Sync(
		t,
		"module alfa\n\ngo 1.27\n\nreplace delta.test/d => ../delta\n",
		go_mod_tester.Replace("bravo.test/b", "v1.0.0"),
		go_mod_tester.Replace("charlie.test/c", "v2.0.0"),
	)
	assert.String(
		t,
		first,
		go_mod_tester.Sync(
			t,
			first,
			go_mod_tester.Replace("bravo.test/b", "v1.0.0"),
			go_mod_tester.Replace("charlie.test/c", "v2.0.0"),
		),
	)
}

func TestSyncReplacesRemovesEmptyManagedBlock(t *testing.T) {
	assert.String(
		t,
		"module alfa\n\ngo 1.27\n\nreplace delta.test/d => ../delta\n",
		go_mod_tester.Sync(
			t,
			"module alfa\n\ngo 1.27\n\nreplace delta.test/d => ../delta\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.0.0\n)\n",
		),
	)
}

func TestSyncReplacesSingleEntryStaysManaged(t *testing.T) {
	once := go_mod_tester.Sync(
		t,
		"module alfa\n",
		go_mod_tester.Replace("bravo.test/b", "v1.0.0"),
	)
	assert.String(
		t,
		"module alfa\n\n// managed by replace\nreplace (\n\tbravo.test/b => bravo.test/b v1.1.0\n)\n",
		go_mod_tester.Sync(
			t,
			once,
			go_mod_tester.Replace("bravo.test/b", "v1.1.0"),
		),
	)
}

func TestSyncReplacesKeepsUpstreamComments(t *testing.T) {
	upstream, e := modfile.Parse(
		constant.ModFile,
		[]byte(
			"module upstream\n\nreplace (\n\t// Avoid a CVE\n\tbravo.test/b => bravo.test/b v1.0.0\n)\n",
		),
		nil,
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"module alfa\n\n// managed by replace\nreplace (\n\t// Avoid a CVE\n\tbravo.test/b => bravo.test/b v1.0.0\n)\n",
		go_mod_tester.Sync(t, "module alfa\n", upstream.Replace...),
	)
}

func TestSyncReplacesInvalid(t *testing.T) {
	_, e := go_mod.SyncReplaces([]byte("module\nreplace =>\n"), nil)
	assert.Error(t, e)
}
