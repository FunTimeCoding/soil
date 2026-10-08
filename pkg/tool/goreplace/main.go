package goreplace

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"io"
	"os"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argumentConstant.DryRun,
		false,
		"Print what would change without writing",
	)
	a.Parse()
	path := a.RequiredPositional(0, "FILE")
	input, e := io.ReadAll(os.Stdin)
	errors.PanicOnError(e)

	if f := Run(
		path,
		string(input),
		a.GetBoolean(argumentConstant.DryRun),
		os.Stdout,
	); f != nil {
		system.Exitf(1, "%s\n", f)
	}
}
