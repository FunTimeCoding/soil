package unit

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/matrix"
)

func frontends(v *virtual_file_system.System) []*matrix.Frontend {
	return scan.Frontends(
		v,
		scan.Services(v, "test", audit_configuration.New()),
	)
}
