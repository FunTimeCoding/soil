package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestFindFilesByExtension(t *testing.T) {
	fixture := join.Absolute(
		system.FindDirectoryUp(system.WorkDirectory(), git.Directory),
		constant.FixturePath,
	)
	assert.Count(
		t,
		5,
		system.FindFilesByExtension(fixture, ".json"),
	)
}
