package gopackageapk

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/semver"
	"github.com/funtimecoding/soil/pkg/tool/gopackageapk/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopackageapk/option"
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
	o := option.New()
	o.Executable = a.RequiredPositional(0, "EXECUTABLE")
	o.PackageVersion = semver.Trim(a.RequiredPositional(1, "PACKAGE_VERSION"))
	Run(o)
}
