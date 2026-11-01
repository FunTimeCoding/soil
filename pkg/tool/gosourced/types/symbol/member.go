package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func (s *Symbol) Member() string {
	if s.Receiver == "" {
		return s.Name
	}

	return join.Empty(s.Receiver, constant.MemberSeparator, s.Name)
}
