package label

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
	"slices"
)

func (s *Service) record(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	old []string,
) (string, error) {
	now, e := s.store.GetLabel(sessionIdentifier, key)

	if e != nil {
		return "", e
	}

	past := join.Comma(old)
	current := join.Comma(now)

	if slices.Equal(old, now) {
		if current == "" {
			return fmt.Sprintf(constant.LabelNotSet, key), nil
		}

		return fmt.Sprintf(constant.LabelUnchanged, key, current), nil
	}

	if f := s.store.LogEvent(
		sessionIdentifier,
		constant.Label,
		name,
		map[string]string{
			constant.Target: target,
			constant.Key:    key,
			constant.Past:   past,
			constant.Now:    current,
		},
	); f != nil {
		return "", f
	}
	s.notify()

	return label_change.Format(key, past, current), nil
}
