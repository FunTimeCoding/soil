package gomcp

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gomcp/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argument.Token, "", "Bearer token for authorization")
	a.Parse(version, gitHash, buildDate)
	probe(
		a.RequiredPositional(0, "URL"),
		a.GetString(argument.Token),
	)
}
