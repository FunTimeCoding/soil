package xref

import (
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"slices"
)

func Extract(
	p *packages.Package,
	workspace map[string]bool,
) *References {
	result := NewReferences()

	for i, o := range p.TypesInfo.Uses {
		if o == nil || o.Pkg() == nil || o.Pkg() == p.Types ||
			!workspace[o.Pkg().Path()] {
			continue
		}

		if _, isPackage := o.(*types.PkgName); isPackage {
			continue
		}

		target, okay := Target(o)

		if !okay {
			continue
		}

		position := p.Fset.Position(i.Pos())
		result.Targets[target] = append(
			result.Targets[target],
			NewSite(
				filepath.Base(position.Filename),
				position.Line,
				position.Column,
			),
		)
	}

	for _, sites := range result.Targets {
		slices.SortFunc(sites, compareSites)
	}

	return result
}
