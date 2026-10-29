package unit

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/google/go-github/v93/github"
)

func workflowRun(
	identifier int64,
	branch string,
	status string,
	conclusion string,
) *run.Run {
	return run.New(
		&github.WorkflowRun{
			ID:         new(identifier),
			Name:       new("release image"),
			HeadBranch: new(branch),
			Status:     new(status),
			Conclusion: new(conclusion),
			Repository: &github.Repository{},
		},
	)
}
