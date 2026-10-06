package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"golang.org/x/tools/go/packages"
	"slices"
)

func constructorHome(
	p *packages.Package,
	name string,
) (string, string, string) {
	if !token.IsExported(name) {
		return "", "", fmt.Sprintf(
			"%s is unexported - no literal outside %s needs a constructor",
			name,
			p.PkgPath,
		)
	}

	structs, receivers := packageStructs(p)
	target := join.Empty(p.PkgPath, "/", camel.ToSnake(name))
	constructor, file := "New", "new.go"

	switch {
	case len(structs) == 1:
	case len(receivers) == 0:
		constructor = join.Empty("New", name)
		file = join.Empty("new_", camel.ToSnake(name), constant.GoExtension)
	case slices.Contains(receivers, name):
		return "", "", fmt.Sprintf(
			"%s carries methods beside other structs in %s - move it out first with extract_type, e.g. to %s",
			name,
			p.PkgPath,
			target,
		)
	default:
		return "", "", fmt.Sprintf(
			"%s shares %s with receiver struct %s - move it out first with extract_type, e.g. to %s",
			name,
			p.PkgPath,
			join.CommaSpace(receivers),
			target,
		)
	}

	if p.Types.Scope().Lookup(constructor) != nil {
		return "", "", fmt.Sprintf(
			"%s already exists in %s",
			constructor,
			p.PkgPath,
		)
	}

	return constructor, file, ""
}
