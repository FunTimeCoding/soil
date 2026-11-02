package gomodule

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system"
)

func updateGitLabFile(goVersion string) {
	if !system.FileExists(soil.GitLabFile) {
		return
	}

	system.SaveFile(
		soil.GitLabFile,
		project.ReplaceGoImageVersion(
			system.ReadFile(system.WorkDirectory(), soil.GitLabFile),
			goVersion,
		),
	)
}
