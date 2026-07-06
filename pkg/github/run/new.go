package run

import (
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/google/go-github/v88/github"
	"time"
)

func New(v *github.WorkflowRun) *Run {
	var create time.Time
	var update time.Time

	if v.CreatedAt != nil {
		create = v.CreatedAt.Time
	}

	if v.UpdatedAt == nil {
		update = create
	} else {
		update = v.UpdatedAt.Time
	}

	return &Run{
		MonitorIdentifier: constant.GoGitHubJob.Integer64Identifier(
			v.GetID(),
		),
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		Conclusion: v.GetConclusion(),
		Branch:     v.GetHeadBranch(),
		Create:     create,
		Update:     update,
		Repository: repository.New(v.Repository),
		Raw:        v,
	}
}
