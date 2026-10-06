package index

import "go/types"

func qualifier(p *types.Package) string {
	return p.Path()
}
