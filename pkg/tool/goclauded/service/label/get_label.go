package label

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Service) GetLabel(
	sessionIdentifier string,
	key string,
) (string, error) {
	values, e := s.store.GetLabel(sessionIdentifier, key)

	return join.Comma(values), e
}
