package repository

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"path/filepath"
	"strings"
)

func New(
	path string,
	status string,
) *Repository {
	name := filepath.Base(path)

	return &Repository{
		MonitorIdentifier: constant.GoGitStatus.StringIdentifier(name),
		Name:              name,
		Path:              path,
		IsClean:           len(strings.TrimSpace(status)) == 0,
		Status:            status,
	}
}
