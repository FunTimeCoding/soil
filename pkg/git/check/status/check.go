package status

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Check(o *option.Status) {
	elements := repository.Filter(
		monitor.OnlyConcerns(collect(o.Path, o.Depth), o.All),
		environment.Slice(constant.RepositoryExcludeEnvironment),
		o.All,
	)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format.Copy()

	if o.Verbose {
		f.Tag(consoleConstant.TagChanges)
	}

	for _, r := range elements {
		console.Line(r.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoGitStatus.Plural)
	}
}
