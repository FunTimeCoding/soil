package commit

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/gocommit/commit/option"
)

func Run(o *option.Commit) {
	validate(o)
	c := gitlab.New(o.Host, o.Token)
	p := common.FindProjectOrExit(c, o.Owner, o.Repository)
	c.MustCommit(
		p.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		strings.ReplaceAllSlice(
			system.ReadFileUnsafe(o.Template),
			o.Replace,
		),
		c.MustFileExists(p, o.Branch, o.Path),
	)
}
