package result

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

type Relations struct {
	Relations []client.Relation `json:"relations"`
	Summary   *RelationSummary  `json:"summary"`
}
