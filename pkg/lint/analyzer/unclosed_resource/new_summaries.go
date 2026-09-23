package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func NewSummaries(loaded []*packages.Package) *Summaries {
	result := &Summaries{arranges: make(map[*types.Func]bool)}
	delegate := make(map[*types.Func][]*types.Func)

	for _, p := range loaded {
		if p.TypesInfo == nil {
			continue
		}

		for _, file := range p.Syntax {
			for _, d := range file.Decls {
				f, okay := d.(*ast.FuncDecl)

				if !okay || f.Body == nil {
					continue
				}

				o, okay := p.TypesInfo.Defs[f.Name].(*types.Func)

				if !okay {
					continue
				}

				if directArranges(p, f.Body) {
					result.arranges[o] = true

					continue
				}

				for _, g := range delegatesOf(p, f.Body) {
					if g != o {
						delegate[o] = append(delegate[o], g)
					}
				}
			}
		}
	}

	for changed := true; changed; {
		changed = false

		for f, all := range delegate {
			if result.arranges[f] {
				continue
			}

			for _, g := range all {
				if result.arranges[g] {
					result.arranges[f] = true
					changed = true

					break
				}
			}
		}
	}

	return result
}
