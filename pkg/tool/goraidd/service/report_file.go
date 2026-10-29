package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (s *Service) reportFile(name string) (string, error) {
	if !filepath.IsLocal(name) || filepath.Base(name) != name {
		return "", not_found.New("report", name)
	}

	path := filepath.Join(s.outputPath, name)

	if !system.FileExists(path) {
		return "", not_found.New("report", name)
	}

	return path, nil
}
