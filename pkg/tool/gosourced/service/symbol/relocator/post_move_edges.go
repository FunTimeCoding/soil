package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strconv"
)

func postMoveEdges(
	all []*packages.Package,
	source *packages.Package,
	entries []*relocation.Entry,
	packagePath string,
	targetPackagePath string,
) map[string]map[string]bool {
	result := make(map[string]map[string]bool)
	seen := make(map[*packages.Package]bool)
	var walk func(*packages.Package)
	walk = func(p *packages.Package) {
		if seen[p] {
			return
		}

		seen[p] = true

		if result[p.PkgPath] == nil {
			result[p.PkgPath] = make(map[string]bool)
		}

		for path, i := range p.Imports {
			result[p.PkgPath][path] = true
			walk(i)
		}
	}

	for _, p := range all {
		walk(p)
	}

	moved := make(map[types.Object]bool)

	for _, entry := range entries {
		moved[entry.Object] = true
	}

	usesMoved := make(map[string]bool)
	keepsSource := make(map[string]bool)

	for _, p := range all {
		if p.PkgPath == packagePath {
			continue
		}

		uses, keeps := sourceUse(p, moved, packagePath)
		usesMoved[p.PkgPath] = usesMoved[p.PkgPath] || uses
		keepsSource[p.PkgPath] = keepsSource[p.PkgPath] || keeps
	}

	for path, uses := range usesMoved {
		if !uses {
			continue
		}

		if path != targetPackagePath {
			result[path][targetPackagePath] = true
		}

		if !keepsSource[path] {
			delete(result[path], packagePath)
		}
	}

	if sourceReferencesMoved(source, entries, moved) {
		result[packagePath][targetPackagePath] = true
	}

	if result[targetPackagePath] == nil {
		result[targetPackagePath] = make(map[string]bool)
	}

	for _, entry := range entries {
		if len(entry.BackIdentifiers) > 0 {
			result[targetPackagePath][packagePath] = true
		}

		for _, c := range entry.Carried {
			if path, f := strconv.Unquote(c.Path.Value); f == nil &&
				path != targetPackagePath {
				result[targetPackagePath][path] = true
			}
		}
	}

	return result
}
