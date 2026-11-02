package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/pipeline_view"
)

func (s *Service) JobView(
	x context.Context,
	instance string,
	project int64,
	pipeline int64,
	requested int64,
) (*pipeline_view.View, error) {
	jobs, e := s.PipelineJobs(x, instance, project, pipeline)

	if e != nil {
		return nil, e
	}

	result := pipeline_view.New(instance, project, pipeline)
	result.Jobs = jobs
	result.Selected = selectJob(jobs, requested)

	if result.Selected == nil {
		return result, nil
	}

	trace, f := s.Trace(x, instance, project, result.Selected.Identifier)

	if f != nil {
		return nil, f
	}

	result.Trace = trace

	return result, nil
}
