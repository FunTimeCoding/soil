package reference

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/location"
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
	references := lookup.ObjectReferences(
		all,
		func(o types.Object) bool {
			return lookup.SameObject(o, declaration)
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
