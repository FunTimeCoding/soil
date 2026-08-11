package scan

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func Frontends(
	v *virtual_file_system.System,
	services []*Service,
) []*Frontend {
	var result []*Frontend

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
