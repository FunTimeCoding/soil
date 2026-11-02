package repository

import (
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"path/filepath"
	"strings"
)

func New(
	path string,
	status string,
) *Repository {
	name := filepath.Base(path)

	return &Repository{
		MonitorIdentifier: monitorConstant.GoGitStatus.StringIdentifier(name),
		Name:              name,
		Path:              path,
		IsClean:           len(strings.TrimSpace(status)) == 0,
		Status:            status,
	}
}
