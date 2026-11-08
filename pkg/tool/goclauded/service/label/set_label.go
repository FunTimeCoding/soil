package label

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/label_value"

func (s *Service) SetLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	value string,
) (string, error) {
	old, e := s.store.SetLabel(sessionIdentifier, key, label_value.Parse(value))

	if e != nil {
		return "", e
	}

	return s.record(sessionIdentifier, name, target, key, old)
}
