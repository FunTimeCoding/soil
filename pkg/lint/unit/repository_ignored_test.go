package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestRepositoryIgnoredReadsEachCheckoutRules(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(
		root,
		"go.mod",
		"module alfa\n\nrequire example.org/bravo v1.0.0\n",
	)
	system_tester.New(t).Write(root, ".gitignore", "/scratch\n")
	system_tester.New(t).Write(
		filepath.Join(root, soil.ParentDirectory, "bravo"),
		".gitignore",
		"/tmp\n",
	)
	repo, _ := lint.Walk(root, option.New("", false))
	assert.Strings(t, []string{"../bravo"}, repo.Siblings)
	assert.True(t, repo.Ignored("scratch/note.md"))
	assert.True(t, repo.Ignored("../bravo/tmp/harness"))
	assert.True(t, repo.Ignored("../bravo/tmp/harness/"))
	assert.Boolean(t, false, repo.Ignored("tmp/harness"))
	assert.Boolean(t, false, repo.Ignored("../bravo/scratch/note.md"))
	assert.Boolean(t, false, repo.Ignored("../bravo/doc/note.md"))
	assert.Boolean(t, false, repo.Ignored("../charlie/tmp/harness"))
}
