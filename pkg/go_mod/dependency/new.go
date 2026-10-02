package dependency

import "github.com/funtimecoding/soil/pkg/monitor/constant"

func New(
	path string,
	version string,
	directory string,
) *Dependency {
	return &Dependency{
		MonitorIdentifier: constant.GoLicense.StringIdentifier(path),
		Path:              path,
		Version:           version,
		Directory:         directory,
	}
}
