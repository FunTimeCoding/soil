package goterraformd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"strings"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Database()
	a.Repository()
	a.CloneDirectory()
	a.CopyModes()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Metric.Address = a.MetricAddress()
	o.Repository = a.Required(argumentConstant.Repository)
	o.CloneDirectory = a.Required(argumentConstant.CloneDirectory)
	o.CopyModes = a.GetBoolean(argumentConstant.CopyModes)
	o.TerraformDirectory = environment.Required(
		constant.TerraformDirectoryEnvironment,
	)
	o.Hostname = system.Hostname()

	if system.FileExists(kubernetesConstant.NamespacePath) {
		o.Namespace = strings.TrimSpace(
			system.ReadFileUnsafe(kubernetesConstant.NamespacePath),
		)
	}

	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.StateNamespace = environment.Fallback(
		constant.StateNamespaceEnvironment,
		constant.StateNamespace,
	)
	o.StateLease = environment.Fallback(
		constant.StateLeaseEnvironment,
		constant.StateLease,
	)

	if v := environment.Fallback(constant.DownstreamEnvironment, ""); v != "" {
		o.Downstream = split.Comma(v)
		o.DownstreamToken = environment.Required(
			constant.DownstreamTokenEnvironment,
		)
	}

	Run(o, s)
}
