package unit

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func frontends(v *virtual_file_system.System) []*scan.Frontend {
	return scan.Frontends(v, scan.Services(v, "test", scan.NewConfiguration()))
}
