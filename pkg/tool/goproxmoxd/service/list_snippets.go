package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"io/fs"
)

func (s *Service) ListSnippets(instance string) ([]fs.FileInfo, error) {
	c, e := s.shells.Client(instance)

	if e != nil {
		return nil, e
	}

	entries, e := c.ListDirectory(constant.SnippetDirectory)

	if e != nil {
		return nil, s.forgetDropped(instance, c, e)
	}

	var result []fs.FileInfo

	for _, v := range entries {
		if v.IsDir() {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
