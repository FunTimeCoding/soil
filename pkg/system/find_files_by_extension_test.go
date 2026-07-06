package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestFindFilesByExtension(t *testing.T) {
	fixture := join.Absolute(
		FindDirectoryUp(WorkDirectory(), git.Directory),
		constant.FixturePath,
	)
	assert.Count(
		t,
		5,
		FindFilesByExtension(fixture, ".json"),
	)
}
