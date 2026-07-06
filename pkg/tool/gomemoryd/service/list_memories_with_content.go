package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListMemoriesWithContent(tag string) ([]store.Memory, error) {
	summaries, e := s.store.ListMemories("", tag, true)

	if e != nil {
		return nil, e
	}

	var result []store.Memory

	for _, sum := range summaries {
		m, e := s.store.GetMemory(sum.Identifier)

		if e != nil {
			continue
		}

		result = append(result, *m)
	}

	return result, nil
}
