package gomonitord

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/monitor/coder"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/example_client"
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

	if false {
		// Start server: go run cmd/gomonitord/main.go localhost:3002
		// Connect: websocat ws://127.0.0.1:3002 --protocol echo
		errors.PanicOnError(coder.Run())
	}

	if false {
		example_client.Run()
	}

	if true {
		gorilla.Run(
			web.AddressHostPort(
				webConstant.Localhost,
				webConstant.ListenPort,
			),
		)
	}
}
