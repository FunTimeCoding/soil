package label

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) SwapLabel(
	sessionIdentifier string,
	name string,
	key string,
	expect string,
	value string,
) (string, error) {
	old, e := s.store.GetLabel(sessionIdentifier, key)

	if e != nil {
		return "", e
	}

	applied, f := s.store.SwapLabel(sessionIdentifier, key, expect, value)

	if f != nil {
		return "", f
	}

	if !applied && expect == "" {
		return "", conflict.Format(constant.LabelHeld, key)
	}

	if !applied {
		return "", conflict.Format(constant.LabelNotHeld, key, expect)
	}

	return s.record(sessionIdentifier, name, "", key, old)
}
