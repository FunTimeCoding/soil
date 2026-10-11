package gomodule

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system"
)

func updateGitLabFile(goVersion string) {
	if !system.FileExists(constant.GitLabFile) {
		return
	}

	system.SaveFile(
		constant.GitLabFile,
		project.ReplaceGoImageVersion(
			system.ReadFile(system.WorkDirectory(), constant.GitLabFile),
			goVersion,
		),
	)
}
