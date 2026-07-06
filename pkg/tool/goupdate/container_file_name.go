package goupdate

import (
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system"
)

func ContainerFileName() string {
	if system.FileExists(project.ContainerFile) {
		return project.ContainerFile
	}

	if system.FileExists(project.DockerFile) {
		return project.DockerFile
	}

	return ""
}
