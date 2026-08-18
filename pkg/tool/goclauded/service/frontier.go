package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"sort"
)

func frontier(
	edges []client.Relation,
	loaded map[int64]bool,
) []door {
	seen := map[int64]*door{}

	for _, edge := range edges {
		add(
			seen,
			loaded,
			edge.TargetIdentifier,
			edge.TargetName,
			&edge,
			edge.SourceIdentifier,
			edge.SourceName,
		)
		add(
			seen,
			loaded,
			edge.SourceIdentifier,
			edge.SourceName,
			&edge,
			edge.TargetIdentifier,
			edge.TargetName,
		)
	}

	var result []door

	for _, one := range seen {
		result = append(result, *one)
	}

	sort.Slice(
		result,
		func(a int, b int) bool {
			return result[a].Identifier < result[b].Identifier
		},
	)

	return result
}
