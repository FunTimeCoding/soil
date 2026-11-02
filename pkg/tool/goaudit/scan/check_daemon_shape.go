package scan

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
	"strings"
)

func (s *Service) checkDaemonShape(
	v *virtual_file_system.System,
	path string,
) {
	s.checkServiceLayout(v, path)
	f := parseWebFile(v, filepath.Join(path, "run.go"))

	if f == nil {
		return
	}

	for _, name := range runConstructors(f) {
		if owner, _, _ := strings.Cut(
			name,
			stringsConstant.Dot,
		); owner == constant.WorkerPackage {
			s.addConcern(
				constant.RunBuildsClockKey,
				fmt.Sprintf(constant.RunBuildsClockText, name),
				path,
			)

			continue
		}

		s.addConcern(
			constant.RunBuildsClientKey,
			formatRunBuildsClient(name),
			path,
		)
	}

	if v.Has(
		filepath.Join(path, constant.DaemonDirectory, constant.DaemonCloseFile),
	) && callsPackage(f, constant.DaemonPackage) && !defersClose(f) {
		s.addConcern(
			constant.DaemonNotClosedKey,
			constant.DaemonNotClosedText,
			path,
		)
	}
}
