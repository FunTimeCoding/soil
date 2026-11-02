package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type PipelineForge interface {
	DeletePipeline(
		x context.Context,
		project int64,
		identifier int64,
	) error
	MustDeletePipeline(
		x context.Context,
		project int64,
		identifier int64,
	)
	Pipeline(
		x context.Context,
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	Pipelines(
		x context.Context,
		project int64,
		reference string,
		status string,
		limit int64,
	) ([]*pipeline.Pipeline, error)
	MustPipelines(
		x context.Context,
		project int64,
	) []*pipeline.Pipeline
	CreatePipeline(
		x context.Context,
		project int64,
		reference string,
		v []*gitlab.PipelineVariableOptions,
	) (*pipeline_detail.Detail, error)
	CancelPipeline(
		x context.Context,
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	RetryPipeline(
		x context.Context,
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
}
