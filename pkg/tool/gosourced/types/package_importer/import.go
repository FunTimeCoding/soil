package package_importer

import (
	"fmt"
	"go/types"
)

func (i *Importer) Import(path string) (*types.Package, error) {
	if p, okay := i.packages[path]; okay {
		return p, nil
	}

	return nil, fmt.Errorf("package %s is not in the loaded build graph", path)
}
