package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"sort"
)

func referenceLocations(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	declaration types.Object,
	excludeFile string,
) []*result.Location {
	var locations []*result.Location

	for _, f := range resolve.FindAllReferences(all, declaration) {
		if f.Ident.Pos() == declaration.Pos() {
			continue
		}

		position := set.Position(f.Ident.Pos())

		if excludeFile != "" && position.Filename == excludeFile {
			continue
		}

		locations = append(
			locations,
			result.NewLocation(
				system.RelativePath(directory, position.Filename),
				position.Line,
				f.Package.PkgPath,
			),
		)
	}

	sort.Slice(
		locations,
		func(i int, j int) bool {
			if locations[i].File != locations[j].File {
				return locations[i].File < locations[j].File
			}

			return locations[i].Line < locations[j].Line
		},
	)

	return locations
}
