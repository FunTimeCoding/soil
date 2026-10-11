package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkGuardBattery(
	v *virtual_file_system.System,
	path string,
) {
	if !s.guarded() {
		return
	}

	if v.Has(
		filepath.Join(
			path,
			constant.UnitDirectory,
			constant.GuardTestFile,
		),
	) || v.Has(
		filepath.Join(
			path,
			constant.IntegrationDirectory,
			"guard",
			constant.GuardTestFile,
		),
	) {
		return
	}

	s.addConcern(
		constant.MissingGuardTestKey,
		constant.MissingGuardTestText,
		path,
	)
}
