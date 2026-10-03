package gomonitord

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse(version, gitHash, buildDate)
	gorilla.Run(
		web.AddressHostPort(webConstant.Localhost, webConstant.ListenPort),
	)
}
