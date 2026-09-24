package repository_tester

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	root := t.TempDir()
	upstream := filepath.Join(root, "upstream.git")
	result := &Tester{Clone: filepath.Join(root, constant.Clone)}
	Command(root, "init", "--bare", "--initial-branch=main", upstream)
	Command(root, constant.Clone, upstream, result.Clone)
	Command(
		result.Clone,
		constant.Configuration,
		"user.email",
		"test@example.com",
	)
	Command(result.Clone, constant.Configuration, "user.name", "Test")
	Command(
		result.Clone,
		constant.Configuration,
		constant.HooksPathKey,
		".git/hooks",
	)
	result.Write(library.ReadmeFile, "# base\n")
	result.Commit("base")
	Command(
		result.Clone,
		constant.Push,
		"-u",
		constant.OriginRemote,
		constant.MainBranch,
	)

	return result
}
