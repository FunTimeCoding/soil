package installed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func tagCommit(
	directory string,
	tag string,
) string {
	r := run.New()
	r.Panic = false
	r.Directory = directory
	output := r.Start(
		constant.Command,
		constant.RevParse,
		join.Empty(tag, lintConstant.CommitSuffix),
	)

	if r.Error != nil {
		return ""
	}

	return strings.TrimSpace(output)
}
