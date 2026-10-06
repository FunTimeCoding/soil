package service

import (
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func indexedLocations(
	directory string,
	i *xref.Index,
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	declaration types.Object,
	excludeFile string,
) []*location.Location {
	var own []*location.Location

	if i.Has(packagePath) {
		own = referenceLocations(
			directory,
			slices.DeleteFunc(
				slices.Clone(all),
				func(p *packages.Package) bool { return p.PkgPath != packagePath },
			),
			set,
			declaration,
			excludeFile,
		)
	}

	return sortLocations(
		append(own, indexLocations(directory, i, declaration, excludeFile)...),
	)
}
