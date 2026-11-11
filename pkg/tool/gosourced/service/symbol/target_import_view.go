package symbol

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"sort"
	"strings"
)

func targetImportView(
	plan *relocation.Plan,
	file *dst.File,
	existing *ast.File,
	entries []*relocation.Entry,
) *import_view.View {
	locals := make(map[string]string)

	for _, entry := range entries {
		for _, spec := range entry.Carried {
			if spec.Name != nil {
				locals[strings.Trim(spec.Path.Value, `"`)] = spec.Name.Name
			}
		}
	}

	for _, spec := range file.Imports {
		path := strings.Trim(spec.Path.Value, `"`)
		delete(locals, path)

		if spec.Name != nil {
			locals[path] = spec.Name.Name
		}
	}

	referenced := make(map[string]bool)
	dst.Inspect(
		file,
		func(n dst.Node) bool {
			if i, okay := n.(*dst.Ident); okay &&
				i.Path != "" &&
				i.Path != plan.TargetPackagePath {
				referenced[i.Path] = true
			}

			return true
		},
	)
	var paths []string

	for path := range referenced {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	var result []*import_entry.Entry

	for _, path := range paths {
		p := pathPackage(plan, path)

		if p == nil {
			continue
		}

		local, aliased := locals[path]

		if !aliased {
			local = p.Name
		}

		modulePath := ""

		if p.Module != nil {
			modulePath = p.Module.Path
		}

		result = append(
			result,
			import_entry.New(path, p.Name, local, modulePath, p.Module == nil),
		)
	}

	v := import_view.New(
		plan.TargetPackagePath,
		result,
		func(name string) bool {
			return targetDeclares(plan, existing, entries, name)
		},
	)

	if m := plan.Source.Module; m != nil {
		own := m.Path
		v.Ours = func(modulePath string) bool { return modulePath == own }
	}

	return v
}
