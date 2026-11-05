package result

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func NewRelations(relations []client.Relation) *Relations {
	return &Relations{
		Relations: relations,
		Summary:   &RelationSummary{Count: len(relations)},
	}
}
