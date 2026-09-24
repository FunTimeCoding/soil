package gohook

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
)

func List(c *configuration.Configuration) {
	console.Line(c.Path)

	for _, name := range c.HookNames() {
		console.Format(
			constant.ListHook,
			name,
			changed.NewHook(c.Root, name).String(),
		)

		for _, j := range c.Hooks[name] {
			if len(j.Paths) == 0 {
				console.Format(constant.ListJob, j.Run)

				continue
			}

			console.Format(
				constant.ListJobPaths,
				j.Run,
				join.CommaSpace(j.Paths),
			)
		}
	}
}
