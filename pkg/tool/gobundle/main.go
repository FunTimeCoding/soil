package gobundle

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gobundle/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobundle/option"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	o := option.New()
	o.Name = a.RequiredPositional(0, "NAME")
	o.Path = a.RequiredPositional(1, "PATH")
	o.Executable = a.RequiredPositional(2, "EXECUTABLE")
	o.Icon = a.RequiredPositional(3, "ICON")
	o.Vendor = a.RequiredPositional(4, "VENDOR")
	o.BundleVersion = a.RequiredPositional(5, "VERSION")
	Run(o)
}
