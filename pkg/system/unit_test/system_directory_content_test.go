package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
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
		system.DirectoryContent(
			join.Absolute(
				system.FindDirectoryUp(system.WorkDirectory(), git.Directory),
				constant.FixturePath,
			),
		),
	)
}
