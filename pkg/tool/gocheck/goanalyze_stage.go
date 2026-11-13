package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/configuration"
	goanalyzeConstant "github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	goanalyzeOption "github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func goanalyzeStage(
	o *option.Check,
	patterns []string,
	shown []string,
	reported map[string]bool,
	handle *index.Workspace,
	w *workspace.Workspace,
) *output.Results {
	header := lint.HeaderLine(
		goanalyzeConstant.Identity.Name(),
		o.Root,
		lint.PatternDetail(shown),
	)
	a := goanalyzeOption.New()
	a.Root = o.Root
	a.Patterns = patterns
	a.Index = o.Index
	a.Comment = configuration.Load(o.Root).Comment
	loaded, okay := gofix.RememberedLoad(w, o.Root, patterns)

	if !okay {
		raw, _, e := resolve.LoadPackagesTagged(
			o.Root,
			handle.Graph().Tags,
			w.Layer().Overlay(),
			patterns...,
		)
		errors.PanicOnError(e)

		if failures := loadFailures(o.Root, raw); len(failures.Entries) > 0 {
			failures.Header = header

			return failures
		}

		loaded = raw
	}

	r := goanalyze.AnalyzeLoaded(a, loaded, reported, handle)
	r.Header = header

	if !o.Fix {
		translate(r, o.Root, w.Layer().Changes())
	}

	return r
}
