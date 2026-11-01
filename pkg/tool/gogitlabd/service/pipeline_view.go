package service

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/pipeline_view"

func (s *Service) PipelineView(
	instance string,
	project int64,
	pipeline int64,
	requested int64,
) (*pipeline_view.View, error) {
	result, e := s.JobView(instance, project, pipeline, requested)

	if e != nil {
		return nil, e
	}

	detail, f := s.Pipeline(instance, project, pipeline)

	if f != nil {
		return nil, f
	}

	result.Detail = detail

	return result, nil
}
