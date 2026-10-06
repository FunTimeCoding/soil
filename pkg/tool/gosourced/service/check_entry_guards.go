package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strconv"
)

func checkEntryGuards(
	all []*packages.Package,
	p *packages.Package,
	target *packages.Package,
	entries []*relocation.Entry,
	packagePath string,
	targetPackagePath string,
	qualifyBackReferences bool,
) string {
	if target != nil && importsTransitively(target, packagePath) {
		return fmt.Sprintf(
			"move would create an import cycle: %s imports %s",
			targetPackagePath,
			packagePath,
		)
	}

	excluded := make(map[token.Pos]bool)

	for _, entry := range entries {
		excluded[entry.Object.Pos()] = true
	}

	for _, entry := range entries {
		if !qualifyBackReferences {
			dependencies := moveDependencies(p, excluded, entry.Node)

			if len(dependencies) > 0 {
				return fmt.Sprintf(
					"%s references package-local symbols: %s",
					entry.Symbol,
					join.CommaSpace(dependencies),
				)
			}
		}

		for _, c := range entry.Carried {
			importPath, f := strconv.Unquote(c.Path.Value)

			if f != nil {
				continue
			}

			if importPath == targetPackagePath {
				continue
			}

			carriedPackage := findPackage(all, importPath)

			if target != nil && carriedPackage != nil &&
				importsTransitively(carriedPackage, targetPackagePath) {
				return fmt.Sprintf(
					"move would create an import cycle through %s",
					importPath,
				)
			}
		}
	}

	return ""
}
