package goatlasd

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	outpostClient "github.com/funtimecoding/soil/pkg/tool/gooutpostd/client"
	gooutpostdConstant "github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
	opnsenseClient "github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	gopnsensedConstant "github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	processClient "github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func peers(o *option.Atlas) {
	if environment.Optional(gopnsensedConstant.TokenEnvironment) != "" {
		o.Lease = func() opnsense.Client {
			return opnsenseClient.NewEnvironment()
		}
	}

	place := environment.Optional(constant.PlaceEnvironment)

	if place != "" &&
		environment.Optional(goprocessdConstant.TokenEnvironment) != "" {
		o.ProcessPlace = place
		o.Process = func() process.Client {
			return processClient.NewEnvironment()
		}
	}

	hosts := environment.Optional(constant.OutpostHostsEnvironment)
	token := environment.Optional(gooutpostdConstant.TokenEnvironment)

	if hosts == "" || token == "" {
		return
	}

	port := environment.RequiredInteger(gooutpostdConstant.PortEnvironment)
	insecure := environment.Exists(gooutpostdConstant.InsecureEnvironment)
	o.OutpostHosts = split.Comma(hosts)
	o.Outpost = func(host string) outpost.Client {
		u := locator.New(host)

		if port != 0 {
			u.Port(port)
		}

		if insecure {
			u.Insecure()
		}

		return outpostClient.New(u, token)
	}
}
