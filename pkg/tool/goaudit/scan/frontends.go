package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/matrix"
)

func Frontends(
	v *virtual_file_system.System,
	services []*Service,
) []*matrix.Frontend {
	var result []*matrix.Frontend

	for _, s := range services {
		if !s.Web {
			continue
		}

		f := scanFrontend(v, s)

		if f != nil {
			result = append(result, f)
		}
	}

	return result
}
