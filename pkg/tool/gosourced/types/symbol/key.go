package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strings"
)

func (s *Symbol) Key() string {
	parts := []string{s.Package, s.Clause, s.Member()}

	if s.Variant != "" {
		parts = append(parts, s.Variant)
	}

	return strings.Join(parts, constant.KeySeparator)
}
