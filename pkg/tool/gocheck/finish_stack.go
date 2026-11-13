package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
)

func finishStack(
	o *option.Check,
	top *layer.Layer,
	result *check_result.Result,
) {
	changes := top.Changes()

	if len(changes) == 0 {
		return
	}

	if !o.Fix {
		result.Changes = changes

		return
	}

	var graphs []*module_graph.Graph

	for _, root := range top.Roots() {
		graphs = append(graphs, module_graph.Build(root))
	}

	o.CommitLock.Lock()
	defer o.CommitLock.Unlock()
	top.BeforeCommit(o.BeforeCommit)
	refused, e := top.Commit(graphs)
	errors.PanicOnError(e)

	if len(refused) == 0 {
		return
	}

	r := output.NewResultsWithDirectory(o.Root)
	r.Header = lint.HeaderLine(constant.Identity.Name(), o.Root, "")

	for _, path := range refused {
		r.AddConcern(
			concern.NewFile(
				lintConstant.ConcurrentWriteKey,
				lintConstant.ConcurrentWriteText,
				path,
				false,
			),
		)
	}

	result.Sections = append(result.Sections, r)
}
