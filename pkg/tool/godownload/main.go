package godownload

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/godownload/constant"
	"github.com/funtimecoding/soil/pkg/tool/godownload/download"
	"github.com/funtimecoding/soil/pkg/tool/godownload/download/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	common.Arguments(a)
	a.String(
		argumentConstant.PackageVersion,
		library.LatestVersion,
		"Version to download, falls back to latest if not found",
	)
	a.String(
		argumentConstant.Output,
		environment.Fallback(
			"OUTPUT",
			constant.DefaultOutput,
		),
		"Output directory for executable",
	)
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	common.ValidateArguments(a)
	o := option.New()
	o.Host = a.GetString(argumentConstant.Host)
	o.Token = a.GetString(argumentConstant.Token)
	o.Owner = a.GetString(argumentConstant.Owner)
	o.Repository = a.GetString(argumentConstant.Repository)
	o.PackageVersion = a.GetString(argumentConstant.PackageVersion)
	o.Output = a.GetString(argumentConstant.Output)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Package = a.RequiredPositional(0, "PACKAGE")
	download.Run(o)
}
