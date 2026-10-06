package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
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
) []*location.Location {
	var locations []*location.Location
	references := objectReferences(
		all,
		func(o types.Object) bool {
			return sameObject(o, declaration)
		},
	)

	for _, f := range references {
		if f.Ident.Pos() == declaration.Pos() {
			continue
		}

		position := set.Position(f.Ident.Pos())

		if excludeFile != "" && position.Filename == excludeFile {
			continue
		}

		locations = append(
			locations,
			location.New(
				system.RelativePath(directory, position.Filename),
				position.Line,
				f.Package.PkgPath,
			),
		)
	}

	sort.Slice(
		locations,
		func(
			i int,
			j int,
		) bool {
			if locations[i].File != locations[j].File {
				return locations[i].File < locations[j].File
			}

			return locations[i].Line < locations[j].Line
		},
	)

	return locations
}
