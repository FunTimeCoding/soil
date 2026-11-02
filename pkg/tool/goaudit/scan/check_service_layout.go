package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
	"strings"
)

func (s *Service) checkServiceLayout(
	v *virtual_file_system.System,
	path string,
) {
	directory := filepath.Join(path, constant.ServiceDirectory)

	if !v.DirectoryExists(directory) {
		s.addConcern(
			constant.MissingServiceKey,
			constant.MissingServiceText,
			path,
		)

		return
	}

	if services := subServices(v, directory); len(services) > 0 &&
		len(sourceFiles(v, directory)) > 0 {
		s.addConcern(
			constant.MixedServiceKey,
			fmt.Sprintf(constant.MixedServiceText, strings.Join(services, ", ")),
			path,
		)
	}

	for _, d := range importingDirectories(
		v,
		directory,
		func(i string) bool { return i == constant.MapperImport },
	) {
		s.addConcern(
			constant.ServiceImportsMapperKey,
			fmt.Sprintf(constant.ServiceImportsMapperText, d),
			path,
		)
	}
}
