package gocheck

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	gofixConstant "github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	gofixOption "github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func gofixStage(
	o *option.Check,
	patterns []string,
	shown []string,
	reported map[string]bool,
	handle *index.Workspace,
	loaded []*packages.Package,
	set *token.FileSet,
	parent *layer.Layer,
) (*workspace.Workspace, *output.Results) {
	f := gofixOption.New()
	f.Root = o.Root
	f.Patterns = patterns
	f.Diff = !o.Fix
	f.Index = o.Index
	f.Reported = reported
	f.Handle = handle
	f.Loaded = loaded
	f.FileSet = set
	f.Layer = parent
	i := o.ModuleInventory()
	f.Replacing = i.Replacing(o.Root)
	f.Ours = i.ModulePaths()
	r := output.NewResultsWithDirectory(o.Root)
	r.Header = lint.HeaderLine(
		gofixConstant.Identity.Name(),
		o.Root,
		lint.PatternDetail(shown),
	)
	w, e := gofix.RunDefault(f, r)

	if e == nil {
		return w, r
	}

	failure, okay := load_failure.As(e)

	if !okay {
		errors.PanicOnError(e)
	}

	r.AddConcern(
		concern.NewFile(
			constant.LoadFailure,
			fmt.Sprintf(
				"does not load: %s - %v",
				failure.Package,
				failure.Cause,
			),
			failure.Package,
			false,
		),
	)

	return nil, r
}
