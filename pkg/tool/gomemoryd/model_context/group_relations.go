package model_context

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func groupRelations(
	edges []store.RelationOverview,
	members map[int64]bool,
) []groupRelation {
	var outward []groupRelation
	var internal []groupRelation

	for _, edge := range edges {
		row := groupRelation{
			SourceIdentifier: edge.SourceIdentifier,
			Source:           edge.SourceName,
			SourceScope:      edge.SourceScope,
			Type:             edge.Type,
			TargetIdentifier: edge.TargetIdentifier,
			Target:           edge.TargetName,
			TargetScope:      edge.TargetScope,
		}

		if members[edge.SourceIdentifier] &&
			members[edge.TargetIdentifier] {
			internal = append(internal, row)

			continue
		}

		outward = append(outward, row)
	}

	return append(outward, internal...)
}
