package overview

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"

func (s *Service) Flagged(name string) (*overview_result.Row, error) {
	f, e := s.flag.Flagged(name)

	if e != nil {
		return nil, e
	}

	return s.row(f), nil
}
