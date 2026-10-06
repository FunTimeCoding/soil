package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func fieldFlips(
	all []*packages.Package,
	entries []*relocation.Entry,
	typeObject types.Object,
	memberNames map[string]bool,
) ([]*relocation.FieldFlip, string) {
	structure, okay := typeObject.Type().(*types.Named).Underlying().(*types.Struct)

	if !okay {
		return nil, ""
	}

	var result []*relocation.FieldFlip

	for i := range structure.NumFields() {
		field := structure.Field(i)

		if field.Exported() {
			continue
		}

		references := resolve.FindAllReferences(all, field)
		outside := false

		for _, f := range references {
			if !insideMoved(entries, f.Ident.Pos()) {
				outside = true

				break
			}
		}

		if !outside {
			continue
		}

		if field.Embedded() {
			return nil, fmt.Sprintf(
				"embedded field %s is referenced outside the type - export it manually first",
				field.Name(),
			)
		}

		newName := FlipName(field.Name())

		if memberNames[newName] {
			return nil, fmt.Sprintf(
				"exporting field %s collides with %s",
				field.Name(),
				newName,
			)
		}

		result = append(
			result,
			relocation.NewFieldFlip(field, newName, references))
	}

	return result, ""
}
