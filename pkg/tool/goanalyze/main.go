package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("summary", false, "One line per file")
	a.Boolean("comment", false, "Include the stray_comment lint")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Summary = a.GetBoolean("summary")
	o.Comment = a.GetBoolean("comment")
	o.Patterns = a.Positionals()
	Run(o)
}
