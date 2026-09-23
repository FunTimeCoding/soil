package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func newPipeline(
	identifier int64,
	reference string,
	status string,
) *pipeline.Pipeline {
	return pipeline.New(
		&gitlab.PipelineInfo{
			ID:     identifier,
			Ref:    reference,
			Status: status,
		},
	)
}
