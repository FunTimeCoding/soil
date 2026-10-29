package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func (s *Service) GetDocument(path string) (*result.Document, []string, error) {
	document, found, e := s.store.FindDocument(path)

	if e != nil {
		return nil, nil, e
	}

	if found {
		return document, nil, nil
	}

	similar, f := s.store.FindSimilarFiles(path, 5)

	if f != nil {
		return nil, nil, f
	}

	return nil, similar, not_found.New(constant.DocumentKind, path)
}
