package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strconv"
)

func (s *Symbol) Reference() string {
	return join.Empty(
		strconv.Quote(s.Path),
		constant.ReferenceSeparator,
		s.Name,
	)
}
