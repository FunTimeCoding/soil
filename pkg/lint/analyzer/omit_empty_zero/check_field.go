package omit_empty_zero

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"reflect"
	"strings"
)

func checkField(
	p *packages.Package,
	results *output.Results,
	field *ast.Field,
) {
	if field.Tag == nil {
		return
	}

	tag := reflect.StructTag(
		strings.Trim(field.Tag.Value, "`"),
	).Get(constant.NotationKey)

	if !hasOmitEmpty(tag) {
		return
	}

	if !isNumericOrBool(p.TypesInfo.TypeOf(field.Type)) {
		return
	}

	if suppress.IsSuppressed(
		p.Fset,
		p.Syntax,
		field.Pos(),
		"omit_empty_zero",
	) {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"omit_empty_zero",
			fmt.Sprintf(
				"field %s: omitempty stops omitting the zero value under json/v2 - use omitzero",
				fieldName(field),
			),
			p.Fset.Position(field.Pos()).Filename,
			false,
		),
	)
}
