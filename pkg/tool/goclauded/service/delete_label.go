package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
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

	change := label_change.Format(key, old, "")

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		map[string]string{
			constant.Target: target,
			constant.Key:    key,
			constant.Past:   old,
			constant.Now:    "",
		},
	); e != nil {
		return "", e
	}
	s.notify()

	return change, nil
}
