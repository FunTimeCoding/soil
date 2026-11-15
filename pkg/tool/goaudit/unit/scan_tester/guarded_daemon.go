package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
)

func GuardedDaemon() *virtual_file_system.System {
	result := Daemon(constant.FixtureClosedRun)
	result.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n",
	)

	return result
}
