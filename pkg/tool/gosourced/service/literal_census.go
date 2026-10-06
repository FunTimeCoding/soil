package service

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func literalCensus(
	all []*packages.Package,
	set *token.FileSet,
	named *types.Named,
	structure *types.Struct,
	packagePath string,
) *literal.Census {
	result := literal.NewCensus()
	seen := map[string]bool{}

	for _, p := range all {
		for _, file := range p.Syntax {
			ranges := assert_call.ArgumentRanges(p, file)
			ast.Inspect(
				file,
				func(n ast.Node) bool {
					target := literalTarget(p.TypesInfo, named, n)

					if target == nil {
						return true
					}

					key := set.Position(target.Pos()).String()

					if seen[key] {
						return true
					}

					seen[key] = true
					result.Total++

					if p.PkgPath == packagePath {
						result.Inside++

						return true
					}

					if inRanges(ranges, target.Pos()) {
						result.Expected++

						return true
					}

					result.Sites = append(
						result.Sites,
						literalSite(p, file, structure, target),
					)

					return true
				},
			)
		}
	}

	return result
}
