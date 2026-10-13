package goffrelay

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/constant"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/relay"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/relay/option"
)

func Main() {
	r := reporter.NewOptional(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argumentConstant.Host,
		firefoxConstant.DefaultHost,
		"Bridge address of gofirefoxd",
	)
	a.String(
		constant.ExtensionFlag,
		constant.Extension,
		"Add-on permitted to spawn the relay",
	)
	a.Parse()
	o := option.New()
	o.Host = a.GetString(argumentConstant.Host)
	o.Extension = a.GetString(constant.ExtensionFlag)
	relay.Run(o, a.Argument(1))
}
