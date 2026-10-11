package fixture

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func Edge(
	source int64,
	sourceName string,
	target int64,
	targetName string,
) client.Relation {
	return client.Relation{
		SourceIdentifier: source,
		SourceName:       sourceName,
		TargetIdentifier: target,
		TargetName:       targetName,
		Type:             new("informs"),
	}
}
