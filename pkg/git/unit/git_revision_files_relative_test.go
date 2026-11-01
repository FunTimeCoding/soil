package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestRevisionFilesAcceptARelativePath(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("alfa.go", "package alfa\n")
	r.Commit("alfa")
	work, e := os.Getwd()
	assert.FatalOnError(t, e)
	relative, f := filepath.Rel(work, r.Clone)
	assert.FatalOnError(t, f)
	files, g := git.RevisionFiles(relative, constant.HeadReference, goFile)
	assert.FatalOnError(t, g)
	assert.Any(t, map[string][]byte{"alfa.go": []byte("package alfa\n")}, files)
}
