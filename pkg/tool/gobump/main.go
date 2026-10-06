package gobump

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gobump/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobump/option"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	o := option.New()
	o.Increase = a.RequiredPositional(0, "INCREASE")
	Run(o)
}
