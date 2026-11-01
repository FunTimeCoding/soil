package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestFindFilesByExtension(t *testing.T) {
	fixture := join.Absolute(
		system.FindDirectoryUp(system.WorkDirectory(), gitConstant.Directory),
		constant.FixturePath,
	)
	assert.Count(t, 10, system.FindFilesByExtension(fixture, ".json"))
}
