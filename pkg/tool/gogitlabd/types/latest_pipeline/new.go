package latest_pipeline

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	branches []*branch.Branch,
	tags []*gitlab.Tag,
	pipelines []*gitlab.PipelineInfo,
) map[string]*gitlab.PipelineInfo {
	branchName := make(map[string]bool)

	for _, b := range branches {
		branchName[b.Name] = true
	}

	tagName := make(map[string]bool)

	for _, t := range tags {
		tagName[t.Name] = true
	}

	result := make(map[string]*gitlab.PipelineInfo)
	var latestTag *gitlab.PipelineInfo

	for _, p := range pipelines {
		if branchName[p.Ref] {
			if e, okay := result[p.Ref]; !okay || p.ID > e.ID {
				result[p.Ref] = p
			}

			continue
		}

		if tagName[p.Ref] &&
			(latestTag == nil || p.ID > latestTag.ID) {
			latestTag = p
		}
	}

	if latestTag != nil {
		result[latestTag.Ref] = latestTag
	}

	return result
}
