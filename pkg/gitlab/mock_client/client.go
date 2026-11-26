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
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type Client struct {
	Closed            int
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
	requests          []*merge_request.Request
	assignedRequests  []*merge_request.Request
	reviewingRequests []*merge_request.Request
	details           []*gitlab.MergeRequest
	notes             []*gitlab.Note
	approvals         map[[2]int64][]string
	threads           map[string][]*gitlab.Note
	resolved          map[string]bool
	registries        []*registry_repository.Repository
	failure           error
}
