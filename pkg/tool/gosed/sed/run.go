package sed

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/strings/base64"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/gosed/sed/option"
	"log"
)

func Run(o *option.Sed) {
	validate(o)
	c := gitlab.New(o.Host, o.Token)
	p := common.FindProjectOrExit(c, o.Owner, o.Repository)

	if len(o.RawActions) > 0 {
		runActions(c, p.Identifier, o)

		return
	}

	f := c.MustFile(p.Identifier, o.Branch, o.Path)

	if f == nil {
		log.Panicf("file does not exist: %s", o.Path)
	}

	c.MustCommit(
		p.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		replaceLines(base64.Decode(f.Content), o.Replaces),
		true,
	)
}
