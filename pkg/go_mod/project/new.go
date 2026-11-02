package project

import (
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"path/filepath"
)

func New(
	path string,
	projectVersion string,
	runtimeVersion string,
) *Project {
	name := filepath.Base(path)

	return &Project{
		MonitorIdentifier: monitorConstant.GoVersion.StringIdentifier(name),
		Name:              name,
		Path:              path,
		Version:           projectVersion,
		runtimeVersion:    runtimeVersion,
	}
}
