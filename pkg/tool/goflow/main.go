package goflow

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentconstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"strings"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Integer(
		argumentconstant.Width,
		constant.Width,
		"columns to wrap prose to",
	)
	a.Boolean(
		argumentconstant.Check,
		false,
		"report what would be rewrapped and change nothing",
	)
	a.Parse(version, gitHash, buildDate)

	if a.ArgumentCount() == 0 {
		system.Exitf(1, "%s\n", constant.Identity.Usage())
	}

	width := a.GetInteger(argumentconstant.Width)
	check := a.GetBoolean(argumentconstant.Check)
	var failures []string
	var touched []string

	for i := 0; i < a.ArgumentCount(); i++ {
		path := a.Argument(i)
		before := system.ReadFileUnsafe(path)
		after, e := Reflow(before, width)

		if e != nil {
			failures = append(
				failures,
				strings.Join([]string{path, e.Error()}, ": "),
			)

			continue
		}

		if after == before {
			continue
		}

		touched = append(touched, path)

		if !check {
			system.WriteFile(path, []byte(after), constant.Mode)
		}
	}

	for _, path := range touched {
		fmt.Println(path)
	}

	if len(failures) > 0 {
		system.Exitf(1, "%s\n", strings.Join(failures, "\n"))
	}

	if check && len(touched) > 0 {
		system.ExitOnCode(1)
	}
}
