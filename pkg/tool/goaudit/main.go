package goaudit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/option"
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
	a.Boolean(
		argumentConstant.Table,
		false,
		"Print compliance tables instead of lint-style concerns",
	)
	a.Boolean(argumentConstant.Web, false, "Print what each web frontend uses")
	a.Parse(version, gitHash, buildDate)
	roots := a.Positionals()

	if len(roots) == 0 {
		root := git.FindDirectory()

		if root == "" {
			system.Exitf(
				1,
				"no repository found above %s, pass a root\n",
				system.WorkDirectory(),
			)
		}

		roots = []string{root}
	}

	for i, root := range roots {
		if !system.DirectoryExists(root) {
			system.Exitf(1, "root does not exist: %s\n", root)
		}

		roots[i] = system.AbsolutePath(root)
	}

	o := option.New()
	o.Roots = roots
	o.Table = a.GetBoolean(argumentConstant.Table)
	o.Web = a.GetBoolean(argumentConstant.Web)
	Run(o)
}
