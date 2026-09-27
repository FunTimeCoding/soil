package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
)

func (s *Service) SetLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	value string,
) (string, error) {
	old, e := s.store.SetLabel(sessionIdentifier, key, value)

	if e != nil {
		return "", e
	}

	if old == value {
		return fmt.Sprintf("%s %s (unchanged)", key, value), nil
	}

	change := label_change.Format(key, old, value)

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		map[string]string{
			constant.Target: target,
			constant.Key:    key,
			constant.Past:   old,
			constant.Now:    value,
		},
	); e != nil {
		return "", e
	}
	s.notify()

	return change, nil
}
