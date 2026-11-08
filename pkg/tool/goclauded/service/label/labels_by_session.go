package label

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Service) LabelsBySession(sessionIdentifier string) ([]label.Label, error) {
	rows, e := s.store.LabelsBySession(sessionIdentifier)

	return group(rows), e
}
