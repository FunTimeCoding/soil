package package_importer

import "go/types"

func New(packages map[string]*types.Package) *Importer {
	return &Importer{packages: packages}
}
