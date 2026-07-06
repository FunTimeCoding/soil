package scan

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func (s *Service) checkServerPatterns(
	v *virtual_file_system.System,
	path string,
) {
	if !s.Generated || !s.Server {
		return
	}

	serverPath := filepath.Join(path, "server")

	for _, name := range v.MustReadDirectory(serverPath) {
		if !strings.HasSuffix(name, constant.GoExtension) {
			continue
		}

		if strings.HasSuffix(name, constant.TestSuffix) {
			continue
		}

		filePath := filepath.Join(serverPath, name)
		content := v.ReadString(filePath)
		f, _, e := parse.Source(name, content)

		if e != nil {
			continue
		}

		s.checkNilNilReturn(f, filePath)
		s.checkHttpError(f, filePath)
	}

	s.checkServerCaptureFail(v, path)
}
