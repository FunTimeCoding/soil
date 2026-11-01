package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"github.com/funtimecoding/soil/pkg/gitlab/types/recorded_commit"
)

type Client struct {
	files             map[string]*file.File
	commits           []*recorded_commit.Commit
	branches          []*branch.Branch
	tags              []*tag.Tag
	projects          []*project.Project
	projectFailure    error
	pipelineFailure   error
	pipelines         []*pipeline.Pipeline
	deletedPipelines  []int64
	jobs              []*job.Job
	traces            map[int64]string
	assignedRequests  []*merge_request.Request
	reviewingRequests []*merge_request.Request
	registries        []*registry_repository.Repository
	failure           error
}
