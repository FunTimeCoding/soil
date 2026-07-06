package project

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"path/filepath"
)

func New(
	path string,
	projectVersion string,
	runtimeVersion string,
) *Project {
	name := filepath.Base(path)

	return &Project{
		MonitorIdentifier: constant.GoVersion.StringIdentifier(name),
		Name:              name,
		Path:              path,
		Version:           projectVersion,
		runtimeVersion:    runtimeVersion,
	}
}
