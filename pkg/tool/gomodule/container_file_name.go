package gomodule

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func containerFileName() string {
	if system.FileExists(constant.ContainerFile) {
		return constant.ContainerFile
	}

	if system.FileExists(constant.DockerFile) {
		return constant.DockerFile
	}

	return ""
}
