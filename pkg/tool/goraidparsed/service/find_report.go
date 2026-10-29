package service

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"strings"
)

func (s *Service) findReport(workdir string) string {
	for _, entry := range system.ReadDirectory(workdir) {
		if strings.HasSuffix(entry.Name(), soil.HypertextExtension) &&
			entry.Name() != filepath.Base(s.templatePath) {
			return filepath.Join(workdir, entry.Name())
		}
	}

	return ""
}
