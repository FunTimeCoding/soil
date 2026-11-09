package gopostgresd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func Arguments(a *argument.Instance) {
	a.Web()
	a.Inventory()
	a.String(
		constant.BackupDirectoryArgument,
		environment.Fallback(constant.BackupDirectoryEnvironment, ""),
		constant.BackupDirectoryUsage,
	)
}
