package gobreak

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Notation, false, "Print the results as JSON")
	a.Parse()
	results, passed, e := Results(
		a.RequiredPositional(constant.BreakFile, constant.BreakFileUsage),
	)

	if e != nil {
		system.Exitf(1, constant.ProblemFormat, e)
	}

	report(results, passed, a.GetBoolean(argumentConstant.Notation))
}
