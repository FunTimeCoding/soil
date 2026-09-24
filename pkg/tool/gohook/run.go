package gohook

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git/changed"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
)

func Run(
	c *configuration.Configuration,
	o *option.Run,
) int {
	jobs := c.Hooks[o.Hook]

	if len(jobs) == 0 {
		if o.Verbose {
			console.Format(constant.NoJobs, o.Hook)
		}

		return 0
	}

	r := ResolveRange(c.Root, o)

	if r.None {
		if o.Verbose {
			console.Format(constant.NothingPushed, o.Hook)
		}

		return 0
	}

	files := r.Files(c.Root)
	guard := o.Hook == git.HookPreCommit

	for i, j := range jobs {
		if !r.All && !j.Matches(files) {
			if o.Verbose {
				console.Format(constant.Skipped, o.Hook, j.Run)
			}

			continue
		}

		command := j.Command(o.Arguments)

		if o.Verbose {
			console.Format(constant.Running, o.Hook, command)
		}

		var before []string

		if guard {
			before = changed.Unstaged(c.Root)
		}

		e := run.New()
		e.Panic = false
		e.Directory = c.Root
		e.Execute(systemConstant.Shell, systemConstant.ShellCommand, command)

		if e.Error != nil {
			console.Format(constant.JobFailed, o.Hook, i+1, command)

			return 1
		}

		if guard {
			if touched := NewEntries(before, changed.Unstaged(c.Root)); len(touched) > 0 {
				console.Format(
					constant.JobModified,
					o.Hook,
					i+1,
					command,
					join.NewLine(touched),
				)

				return 1
			}
		}
	}

	return 0
}
