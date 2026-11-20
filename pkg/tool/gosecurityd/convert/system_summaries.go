package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func SystemSummaries(v []*system_result.Summary) []*server.SystemSummary {
	result := make([]*server.SystemSummary, len(v))

	for i, s := range v {
		result[i] = SystemSummary(s)
	}

	return result
}
