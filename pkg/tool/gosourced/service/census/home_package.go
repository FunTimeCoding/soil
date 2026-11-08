package census

import "golang.org/x/tools/go/packages"

func homePackage(
	loaded []*packages.Package,
	path string,
) *packages.Package {
	for _, p := range loaded {
		if p.PkgPath == path {
			return p
		}
	}

	return nil
}
