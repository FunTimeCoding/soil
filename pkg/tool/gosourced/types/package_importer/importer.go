package package_importer

import "go/types"

type Importer struct {
	packages map[string]*types.Package
}
