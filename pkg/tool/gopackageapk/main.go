package gopackageapk

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/semver"
	"github.com/funtimecoding/soil/pkg/tool/gopackageapk/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopackageapk/option"
)

func Main() {
	r := reporter.New(constant.Identity.Name())
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	o := option.New()
	o.Executable = a.RequiredPositional(0, "EXECUTABLE")
	o.PackageVersion = semver.Trim(a.RequiredPositional(1, "PACKAGE_VERSION"))
	Run(o)
}
