package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func Arguments(a *argument.Instance) {
	a.Web()
	a.Database()
	a.String(
		argumentConstant.Board,
		environment.Fallback(constant.BoardPathEnvironment, constant.BoardFile),
		constant.BoardUsage,
	)
}
