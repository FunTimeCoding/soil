package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type PipelineForge interface {
	DeletePipeline(
		project int64,
		identifier int64,
	) error
	MustDeletePipeline(
		project int64,
		identifier int64,
	)
	Pipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	Pipelines(
		project int64,
		reference string,
		status string,
		limit int64,
	) ([]*pipeline.Pipeline, error)
	MustPipelines(project int64) []*pipeline.Pipeline
	CreatePipeline(
		project int64,
		reference string,
		v []*gitlab.PipelineVariableOptions,
	) (*pipeline_detail.Detail, error)
	CancelPipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	RetryPipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
}
