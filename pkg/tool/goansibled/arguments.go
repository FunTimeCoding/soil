package goansibled

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
)

func Arguments(a *argument.Instance) {
	a.Web()
	a.Metric()
	a.Database()
	a.Repository()
	a.CloneDirectory()
	a.CopyModes()
	a.String(
		argumentConstant.Playbook,
		environment.Fallback(constant.AnsiblePlaybookEnvironment, ""),
		constant.PlaybookUsage,
	)
}
