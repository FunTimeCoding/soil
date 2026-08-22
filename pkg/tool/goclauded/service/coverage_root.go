package service

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func coverageRoot() string {
	if v := environment.Optional(constant.CoverageRootEnvironment); v != "" {
		return v
	}

	return library.CurrentDirectory
}
