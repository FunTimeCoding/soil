package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func SystemDetail(v *system_result.Detail) *server.SystemDetail {
	findings := make([]*server.SystemFinding, len(v.Findings))

	for i, f := range v.Findings {
		findings[i] = SystemFinding(f)
	}

	return &server.SystemDetail{
		Summary:  *SystemSummary(v.Summary),
		Findings: findings,
	}
}
