package response

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func GroupRelations(
	edges []record.RelationOverview,
	members map[int64]bool,
) []GroupRelation {
	var outward []GroupRelation
	var internal []GroupRelation

	for _, edge := range edges {
		row := GroupRelation{
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
