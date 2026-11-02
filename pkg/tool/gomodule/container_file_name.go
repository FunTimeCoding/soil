package gomodule

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func containerFileName() string {
	if system.FileExists(soil.ContainerFile) {
		return soil.ContainerFile
	}

	if system.FileExists(soil.DockerFile) {
		return soil.DockerFile
	}

	return ""
}
