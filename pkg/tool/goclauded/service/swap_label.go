package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
)

func (s *Service) SwapLabel(
	sessionIdentifier string,
	name string,
	key string,
	expect string,
	value string,
) (string, error) {
	applied, e := s.store.SwapLabel(sessionIdentifier, key, expect, value)

	if e != nil {
		return "", e
	}

	if !applied {
		return "", conflict.Format(
			"label %s does not hold %q, nothing written",
			key,
			expect,
		)
	}

	if expect == value {
		return label_change.Format(key, expect, value), nil
	}

	if f := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		map[string]string{
			constant.Key:  key,
			constant.Past: expect,
			constant.Now:  value,
		},
	); f != nil {
		return "", f
	}
	s.notify()

	return label_change.Format(key, expect, value), nil
}
