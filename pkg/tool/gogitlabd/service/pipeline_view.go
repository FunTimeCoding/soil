package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/pipeline_view"
)

func (s *Service) PipelineView(
	x context.Context,
	instance string,
	project int64,
	pipeline int64,
	requested int64,
) (*pipeline_view.View, error) {
	result, e := s.JobView(x, instance, project, pipeline, requested)

	if e != nil {
		return nil, e
	}

	detail, f := s.Pipeline(x, instance, project, pipeline)

	if f != nil {
		return nil, f
	}

	result.Detail = detail

	return result, nil
}
