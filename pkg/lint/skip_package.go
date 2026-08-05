package lint

import "golang.org/x/tools/go/packages"

func SkipPackage(p *packages.Package) bool {
	return len(p.Syntax) == 0 || p.ID != p.PkgPath
}
