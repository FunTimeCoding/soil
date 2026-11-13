package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	lintOption "github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	golintConstant "github.com/funtimecoding/soil/pkg/tool/golint/constant"
)

func golintStage(
	o *option.Check,
	scopes []string,
	fixes *layer.Layer,
) *output.Results {
	l := lintOption.New("", false)
	l.Fix = o.Fix
	l.Layer = fixes
	l.Trees = o.Trees
	l.Documents = o.Documents
	l.Resolved = o.Resolved
	resolved, e := lint.Scopes(o.Root, o.Root, scopes)
	errors.PanicOnError(e)
	l.Scopes = resolved

	return lint.Report(golintConstant.Identity.Name(), o.Root, l)
}
