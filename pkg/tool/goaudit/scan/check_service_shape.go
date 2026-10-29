package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkServiceShape(
	v *virtual_file_system.System,
	path string,
) {
	f := parseWebFile(v, filepath.Join(path, "run.go"))

	if f == nil || !callsPackage(f, constant.LifecyclePackage) {
		return
	}

	if !v.DirectoryExists(filepath.Join(path, constant.ServiceDirectory)) {
		s.addConcern(
			constant.MissingServiceKey,
			constant.MissingServiceText,
			path,
		)
	}

	for _, name := range runConstructors(f) {
		s.addConcern(
			constant.RunBuildsClientKey,
			formatRunBuildsClient(name),
			path,
		)
	}

	if v.Has(
		filepath.Join(path, constant.ServiceDirectory, constant.ServiceCloseFile),
	) && callsPackage(f, constant.ServicePackage) && !defersClose(f) {
		s.addConcern(
			constant.ServiceNotClosedKey,
			constant.ServiceNotClosedText,
			path,
		)
	}
}
