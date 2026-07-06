package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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

	var change string

	switch old {
	case value:
		return fmt.Sprintf("%s %s (unchanged)", key, value), nil
	case "":
		change = fmt.Sprintf("%s (unset)→%s", key, value)
	default:
		change = fmt.Sprintf("%s %s→%s", key, old, value)
	}

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
