package unit

import "go/types"

func qualifiedPath(p *types.Package) string {
	return p.Path()
}
