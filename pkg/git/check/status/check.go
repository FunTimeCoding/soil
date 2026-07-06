package status

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	"github.com/funtimecoding/soil/pkg/monitor"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Check(o *option.Status) {
	elements := repository.Filter(
		monitor.OnlyConcerns(collect(o.Path, o.Depth), o.All),
		environment.Slice(RepositoryExcludeEnvironment),
		o.All,
	)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format.Copy()

	if o.Verbose {
		f.Tag(tag.Changes)
	}

	for _, r := range elements {
		fmt.Println(r.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoGitStatus.Plural)
	}
}
