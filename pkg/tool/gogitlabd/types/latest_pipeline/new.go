package latest_pipeline

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func New(
	branches []*branch.Branch,
	tags []*tag.Tag,
	pipelines []*pipeline.Pipeline,
) map[string]*pipeline.Pipeline {
	branchName := make(map[string]bool)

	for _, b := range branches {
		branchName[b.Name] = true
	}

	tagName := make(map[string]bool)

	for _, t := range tags {
		tagName[t.Name] = true
	}

	result := make(map[string]*pipeline.Pipeline)
	var latestTag *pipeline.Pipeline

	for _, p := range pipelines {
		if branchName[p.Reference] {
			if e, okay := result[p.Reference]; !okay ||
				p.Identifier > e.Identifier {
				result[p.Reference] = p
			}

			continue
		}

		if tagName[p.Reference] &&
			(latestTag == nil || p.Identifier > latestTag.Identifier) {
			latestTag = p
		}
	}

	if latestTag != nil {
		result[latestTag.Reference] = latestTag
	}

	return result
}
