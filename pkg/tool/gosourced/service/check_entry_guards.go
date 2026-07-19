package service

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strconv"
	"strings"
)

func checkEntryGuards(
	all []*packages.Package,
	p *packages.Package,
	target *packages.Package,
	entries []*moveEntry,
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
		excluded[entry.object.Pos()] = true
	}

	for _, entry := range entries {
		if !qualifyBackReferences {
			dependencies := moveDependencies(p, excluded, entry.node)

			if len(dependencies) > 0 {
				return fmt.Sprintf(
					"%s references package-local symbols: %s",
					entry.symbol,
					strings.Join(dependencies, ", "),
				)
			}
		}

		for _, c := range entry.carried {
			importPath, f := strconv.Unquote(c.Path.Value)

			if f != nil {
				continue
			}

			if importPath == targetPackagePath {
				return fmt.Sprintf(
					"%s references the target package",
					entry.symbol,
				)
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
