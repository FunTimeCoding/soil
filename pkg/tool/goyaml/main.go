package goyaml

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goyaml/constant"
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
	a.Parse(version, gitHash, buildDate)

	if a.ArgumentCount() == 0 {
		system.Exitf(1, "%s\n", constant.Identity.Usage())
	}

	var failures []string

	for i := 0; i < a.ArgumentCount(); i++ {
		path := a.Argument(i)

		if e := Check(path); e != nil {
			failures = append(
				failures,
				strings.Join([]string{path, e.Error()}, ": "),
			)
		}
	}

	if len(failures) > 0 {
		system.Exitf(1, "%s\n", strings.Join(failures, "\n"))
	}
}
