package dependency

import monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"

func New(
	path string,
	version string,
	directory string,
) *Dependency {
	return &Dependency{
		MonitorIdentifier: monitorConstant.GoLicense.StringIdentifier(path),
		Path:              path,
		Version:           version,
		Directory:         directory,
	}
}
