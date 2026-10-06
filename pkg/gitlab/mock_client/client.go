package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

type Client struct {
	files             map[string]*file.File
	commits           []*RecordedCommit
	branches          []*branch.Branch
	tags              []*tag.Tag
	projects          []*project.Project
	pipelineFailure   error
	pipelines         []*pipeline.Pipeline
	deletedPipelines  []int64
	assignedRequests  []*merge_request.Request
	reviewingRequests []*merge_request.Request
}
