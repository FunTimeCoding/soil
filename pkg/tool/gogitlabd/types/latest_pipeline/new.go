package latest_pipeline

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	branches []*branch.Branch,
	pipelines []*gitlab.PipelineInfo,
) map[string]*gitlab.PipelineInfo {
	existing := make(map[string]bool)

	for _, b := range branches {
		existing[b.Name] = true
	}

	result := make(map[string]*gitlab.PipelineInfo)

	for _, p := range pipelines {
		if !existing[p.Ref] {
			continue
		}

		if e, okay := result[p.Ref]; !okay || p.ID > e.ID {
			result[p.Ref] = p
		}
	}

	return result
}
