package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"strings"
)

func (s *Service) List(path string) ([]string, error) {
	path = strings.TrimSuffix(path, constant.Folder)

	if path != "" {
		if e := validatePath(path); e != nil {
			return nil, e
		}
	}

	return s.source.List(path)
}
