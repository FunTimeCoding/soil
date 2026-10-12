package imports

import (
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func ViewOf(
	p *packages.Package,
	file *ast.File,
) (*import_view.View, map[*import_entry.Entry]*ast.ImportSpec) {
	var entries []*import_entry.Entry
	specs := make(map[*import_entry.Entry]*ast.ImportSpec)

	for _, spec := range file.Imports {
		if spec.Name != nil && (spec.Name.Name == "_" || spec.Name.Name == ".") {
			continue
		}

		importPath := strings.Trim(spec.Path.Value, `"`)
		imported := ImportedPackage(p, importPath)

		if imported == nil {
			continue
		}

		local := imported.Name()

		if spec.Name != nil {
			local = spec.Name.Name
		}

		e := import_entry.New(
			importPath,
			imported.Name(),
			local,
			ModulePath(p, importPath),
			IsStandardImport(p, importPath),
		)
		entries = append(entries, e)
		specs[e] = spec
	}

	return import_view.New(
		p.PkgPath,
		entries,
		func(name string) bool { return LocalTaken(p, file, name) },
	), specs
}
