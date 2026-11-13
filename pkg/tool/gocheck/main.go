package gocheck

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"os"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argumentConstant.Fix,
		false,
		"Apply what golint and gofix can fix instead of reporting it",
	)
	a.String(
		constant.Reach,
		"",
		"Check only what a session's changes reach: own for this session, or a session's name",
	)
	a.Boolean(
		constant.Fresh,
		false,
		"Check every package again instead of skipping those cached as clean at an unchanged key",
	)
	a.String(
		argumentConstant.Root,
		"",
		"Repository root to check, discovered upward from the working directory when empty",
	)
	a.Boolean(
		argumentConstant.Local,
		false,
		"Run in this process with the lint code on disk instead of asking gosourced",
	)
	a.Parse()
	root, _ := lint.Root(a.GetString(argumentConstant.Root))
	o := option.New()
	o.Root = root
	o.Fix = a.GetBoolean(argumentConstant.Fix)
	o.Fresh = a.GetBoolean(constant.Fresh)

	if reach := a.GetString(constant.Reach); reach != "" {
		transcript, e := Transcript(reach)

		if e != nil {
			system.Exitf(1, "%s\n", e)
		}

		o.Transcript = transcript
	}

	result, e := check(o, a.GetBoolean(argumentConstant.Local))

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	if Print(result) {
		os.Exit(1)
	}
}
