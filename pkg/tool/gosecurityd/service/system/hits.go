package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"

func (s *Service) Hits(
	names []string,
) (map[string]*affected_count.Count, error) {
	details, e := s.Details()

	if e != nil {
		return nil, e
	}

	result := make(map[string]*affected_count.Count, len(names))

	for _, name := range names {
		result[name] = affected_count.New()
	}

	for _, d := range details {
		each := worstEach(d.Findings)

		for _, name := range names {
			if f := each[name]; f != nil {
				result[name].Add(f.Kind())
			}
		}
	}

	return result, nil
}
