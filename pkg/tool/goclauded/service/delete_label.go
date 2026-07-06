package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) DeleteLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
) (string, error) {
	old, e := s.store.DeleteLabel(sessionIdentifier, key)

	if e != nil {
		return "", e
	}

	if old == "" {
		return fmt.Sprintf("%s (not set)", key), nil
	}

	change := fmt.Sprintf("%s %s→ (unset)", key, old)

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		map[string]string{
			constant.Target: target,
			constant.Key:    key,
			"change":        change,
		},
	); e != nil {
		return "", e
	}
	s.notify()

	return change, nil
}
