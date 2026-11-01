package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/pipeline_page"
	"maragu.dev/gomponents"
)

func (s *Server) detailRegion(
	instance string,
	project int64,
	pipeline int64,
	requested int64,
) gomponents.Node {
	v, e := s.service.PipelineView(instance, project, pipeline, requested)
	errors.PanicOnError(e)

	return pipeline_page.New(v).Detail()
}
