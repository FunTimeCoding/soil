package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	git "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
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
