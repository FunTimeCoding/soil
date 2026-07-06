package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func (s *Service) GetDocument(path string) (*result.Document, []string, error) {
	document, e := s.store.GetDocument(path)

	if e != nil {
		return nil, nil, e
	}

	if document != nil {
		return document, nil, nil
	}

	similar, e := s.store.FindSimilarFiles(path, 5)

	if e != nil {
		return nil, nil, e
	}

	return nil, similar, nil
}
