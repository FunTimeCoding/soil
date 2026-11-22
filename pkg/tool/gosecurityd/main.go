package gosecurityd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/option"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.Duration(constant.StaleArgument, constant.StaleAfter, constant.StaleUsage)
	a.Duration(
		constant.TrackerStaleArgument,
		constant.TrackerStaleAfter,
		constant.TrackerStaleUsage,
	)
	a.Parse()
	o := option.New()
	o.Read(a)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.StaleAfter = a.GetDuration(constant.StaleArgument)
	o.TrackerStaleAfter = a.GetDuration(constant.TrackerStaleArgument)

	if environment.Exists(constant.AnsibleHostEnvironment) {
		l := locator.Environment(
			constant.AnsibleHostEnvironment,
			constant.AnsiblePortEnvironment,
			constant.AnsibleInsecureEnvironment,
		)
		token := environment.Required(constant.AnsibleTokenEnvironment)
		o.Ansible = func() face.Client {
			return client.New(l, token)
		}
	}

	Run(o, s)
}
