package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func Snapshot(
	name string,
	v *server.SystemSnapshot,
) (*system.System, []installed.Installed) {
	priority := 0

	if v.Priority != nil {
		priority = *v.Priority
	}

	result := system.New(name, v.Release, priority)

	if v.Kernel != nil {
		result.Kernel = *v.Kernel
	}

	if v.Labels != nil {
		result.Labels = *v.Labels
	}

	packages := make([]installed.Installed, 0, len(v.Packages))

	for _, p := range v.Packages {
		if p == nil {
			continue
		}

		candidate := ""

		if p.Candidate != nil {
			candidate = *p.Candidate
		}

		packages = append(
			packages,
			*installed.New(
				p.Name,
				p.Version,
				p.Source,
				p.SourceVersion,
				candidate,
			),
		)
	}

	return result, packages
}
