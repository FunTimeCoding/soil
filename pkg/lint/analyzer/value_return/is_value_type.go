package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

func isValueType(
	p *packages.Package,
	named *types.Named,
) bool {
	position := p.Fset.Position(named.Obj().Pos())

	if position.Filename == "" || position.Line < 2 {
		return false
	}

	b, e := os.ReadFile(position.Filename)

	if e != nil {
		return false
	}

	lines := strings.Split(string(b), "\n")

	if position.Line > len(lines) {
		return false
	}

	for i := position.Line - 2; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])

		if !strings.HasPrefix(line, constant.CommentPrefix) {
			return false
		}

		if reason, found := strings.CutPrefix(
			line,
			constant.ByValueDirective,
		); found {
			return strings.TrimSpace(reason) != ""
		}
	}

	return false
}
