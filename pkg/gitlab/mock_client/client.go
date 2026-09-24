package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

type Client struct {
	files             map[string]*file.File
	commits           []*RecordedCommit
	branches          []*branch.Branch
	pipelines         []*pipeline.Pipeline
	deletedPipelines  []int64
	assignedRequests  []*merge_request.Request
	reviewingRequests []*merge_request.Request
}
