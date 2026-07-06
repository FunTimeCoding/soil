package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			constant.BoardPath,
			constant.HypertextPath,
			constant.MarkdownPath,
			constant.MemoryPath,
			constant.NotationPath,
			constant.SearchPath,
			constant.WikiPath,
		},
		DirectoryContent(
			join.Absolute(
				FindDirectoryUp(WorkDirectory(), git.Directory),
				constant.FixturePath,
			),
		),
	)
}
