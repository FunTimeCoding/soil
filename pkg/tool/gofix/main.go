package gofix

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"os"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("diff", false, "Show diff without applying")
	a.Boolean("survey", false, "Print violations without fixing")
	a.Boolean("rename", false, "Variable letter rename mode")
	a.Boolean("summary", false, "One line per modified file")
	a.Boolean(
		"full",
		false,
		"Load the whole module for the naming pass instead of the persisted per-package facts",
	)
	a.String(
		argumentConstant.Root,
		"",
		"Repository root to fix, discovered upward from the working directory when empty",
	)
	a.Parse()
	build_tag.Memoize()
	root, work := lint.Root(a.GetString(argumentConstant.Root))
	patterns, e := lint.Patterns(root, work, a.Positionals())

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	lint.Header(constant.Identity.Name(), root, lint.PatternDetail(patterns))

	if a.GetBoolean("survey") {
		runSurvey(root, patterns)

		return
	}

	diff := a.GetBoolean("diff")
	s := output.NewResultsWithDirectory(root)

	if a.GetBoolean("rename") {
		RunVariableNamingFixWithDirectory(patterns, root, diff, s)
	} else {
		o := option.New()
		o.Root = root
		o.Patterns = patterns
		o.Diff = diff
		o.Full = a.GetBoolean("full")
		o.Replacing = inventory.LoadOptional(inventory.DefaultPath()).Replacing(
			root,
		)
		RunDefault(o, s)
	}

	hasBlocked := output.PrintResults(s.Entries, a.GetBoolean("summary"))

	if hasBlocked {
		os.Exit(1)
	}
}
