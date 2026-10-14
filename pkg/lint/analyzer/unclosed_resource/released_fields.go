package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func releasedFields(
	p *packages.Package,
	methods map[string]*ast.FuncDecl,
) []fact.Release {
	result := []fact.Release{}
	visited := make(map[string]bool)
	queue := []string{constant.CloseMethod}

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		f, okay := methods[name]

		if !okay || visited[name] {
			continue
		}

		visited[name] = true
		r := receiverOf(p, f)

		if r == nil {
			continue
		}

		queue = append(queue, receiverCalls(p, f)...)
		result = append(result, releasesIn(p, f.Body, r)...)
	}

	return result
}
