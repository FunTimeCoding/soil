package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func CheckMethodShape(
	p *packages.Package,
	entries []*relocation.Entry,
	target *types.Named,
) string {
	moved := map[string]bool{}

	for _, entry := range entries {
		moved[entry.Symbol] = true
	}

	for _, entry := range entries {
		declaration := entry.Declaration.(*ast.FuncDecl)
		receiver := receiverVariable(p, declaration)

		if receiver == nil {
			continue
		}

		if message := receiverShape(
			p,
			entry.Symbol,
			declaration,
			receiver,
			target,
			moved,
		); message != "" {
			return message
		}
	}

	return ""
}
