package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/receiver_rename"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func planReceivers(
	set *token.FileSet,
	p *packages.Package,
	n *receiver_rename.Rename,
	r *output.Results,
) []*plannedReceiver {
	refuse := func(format string, arguments ...any) []*plannedReceiver {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf(format, arguments...),
				"",
				false,
			),
		)

		return nil
	}

	if !token.IsIdentifier(n.Name) || n.Name == constant.BlankName {
		return refuse("%s is not a receiver name", n.Name)
	}

	if types.Universe.Lookup(n.Name) != nil {
		return refuse("%s is a built-in name", n.Name)
	}

	named, okay := p.Types.Scope().Lookup(n.Type).(*types.TypeName)

	if !okay {
		return refuse("no type %s in %s", n.Type, p.PkgPath)
	}

	var result []*plannedReceiver
	methods := 0

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			method, isFunction := d.(*ast.FuncDecl)

			if !isFunction || receiverTypeName(p, method) != named {
				continue
			}

			methods++
			field := method.Recv.List[0]

			if len(field.Names) == 0 || field.Names[0].Name == constant.BlankName ||
				field.Names[0].Name == n.Name {
				continue
			}

			receiver := p.TypesInfo.Defs[field.Names[0]]
			label := memberName(n.Type, method.Name.Name)

			if taken := receiverCollision(p, method, receiver, n.Name); taken != nil {
				addPositionConcern(
					r,
					set,
					taken.Pos(),
					constant.ConcernValidation,
					false,
					"%s: %s is taken in the method",
					label,
					n.Name,
				)

				continue
			}

			result = append(
				result,
				&plannedReceiver{
					Method: method,
					Idents: receiverUses(p, method, field.Names[0], receiver),
					Label:  label,
					From:   field.Names[0].Name,
					To:     n.Name,
				},
			)
		}
	}

	if methods == 0 {
		return refuse("%s has no methods in %s", n.Type, p.PkgPath)
	}

	return result
}
