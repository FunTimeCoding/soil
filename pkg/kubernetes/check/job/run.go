package job

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
)

func Run() {
	a := argument.NewSimple("kubernetes-job")
	a.Boolean(constant.TrivyArgument, false, "Run Trivy job")
	a.Boolean(constant.LabArgument, false, "Renovate GitLab")
	a.Boolean(constant.HubArgument, false, "Renovate GitHub")
	a.Boolean(argument.Wait, false, "Wait for job to complete")
	a.ParseSimple()
	trivy := a.GetBoolean(constant.TrivyArgument)
	lab := a.GetBoolean(constant.LabArgument)
	hub := a.GetBoolean(constant.HubArgument)
	wait := a.GetBoolean(argument.Wait)
	k := client.NewEnvironment()

	if !trivy && !lab && !hub {
		trivy = true
		lab = true
		hub = true
	}

	if trivy {
		start(
			k,
			wait,
			constant.TrivyNamespace,
			constant.TrivyCron,
			constant.ManualJob,
		)
	}

	if lab {
		start(
			k,
			wait,
			constant.RenovateNamespace,
			constant.LabCron,
			constant.ManualLabJob,
		)
	}

	if hub {
		start(
			k,
			wait,
			constant.RenovateNamespace,
			constant.HubCron,
			constant.ManualHubJob,
		)
	}
}
