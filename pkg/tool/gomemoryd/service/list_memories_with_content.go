package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListMemoriesWithContent(
	tag string,
	scope string,
) ([]record.Memory, error) {
	summaries, e := s.store.ListMemories("", tag, scope, true)

	if e != nil {
		return nil, e
	}

	var result []record.Memory

	for _, sum := range summaries {
		m, e := s.store.GetMemory(sum.Identifier)

		if e != nil {
			continue
		}

		result = append(result, *m)
	}

	return result, nil
}
