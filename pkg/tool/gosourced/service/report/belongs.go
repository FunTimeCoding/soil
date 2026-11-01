package report

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"strings"
)

func belongs(
	s *symbol.Symbol,
	module string,
) bool {
	if s == nil {
		return false
	}

	p := strings.TrimSuffix(s.Path, constant.TestClauseSuffix)

	return p == module ||
		strings.HasPrefix(p, join.Empty(module, stringsConstant.Slash))
}
