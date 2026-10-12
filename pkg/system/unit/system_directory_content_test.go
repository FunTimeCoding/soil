package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"board",
			"claude",
			"directory",
			"hypertext",
			"license",
			"markdown",
			"memory",
			"notation",
			"search",
			"wiki",
		},
		system.DirectoryContent(
			join.Absolute(
				system.FindDirectoryUp(
					system.WorkDirectory(),
					gitConstant.Directory,
				),
				constant.FixturePath,
			),
		),
	)
}
