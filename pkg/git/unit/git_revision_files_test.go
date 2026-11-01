package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"path/filepath"
	"testing"
)

func TestRevisionFilesReadTheCommittedTree(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("alfa.go", "package alfa\n")
	r.Commit("alfa")
	r.Write("alfa.go", "package alfa\n\nvar Bravo = 1\n")
	repository_tester.Command(r.Clone, "add", "alfa.go")
	files, e := git.RevisionFiles(r.Clone, constant.HeadReference, goFile)
	assert.FatalOnError(t, e)
	assert.Any(t, map[string][]byte{"alfa.go": []byte("package alfa\n")}, files)
}

func TestIndexFilesReadTheStagedContent(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("alfa.go", "package alfa\n")
	r.Commit("alfa")
	r.Write("alfa.go", "package alfa\n\nvar Bravo = 1\n")
	repository_tester.Command(r.Clone, "add", "alfa.go")
	r.Write("alfa.go", "package alfa\n\nvar Charlie = 2\n")
	files, e := git.IndexFiles(r.Clone, goFile)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		map[string][]byte{"alfa.go": []byte("package alfa\n\nvar Bravo = 1\n")},
		files,
	)
}

func TestRevisionFilesStayInsideTheirDirectory(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("alfa.go", "package alfa\n")
	r.Write("graft/bravo/bravo.go", "package bravo\n")
	r.Commit("graft")
	files, e := git.RevisionFiles(
		filepath.Join(r.Clone, "graft", "bravo"),
		constant.HeadReference,
		goFile,
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		map[string][]byte{"bravo.go": []byte("package bravo\n")},
		files,
	)
}

func TestRevisionFilesRefuseAnUnknownRevision(t *testing.T) {
	r := repository_tester.New(t)
	_, e := git.RevisionFiles(r.Clone, "nowhere", goFile)
	assert.Error(t, e)
	assert.String(t, "revision not found: nowhere", e.Error())
}
