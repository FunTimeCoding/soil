package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	project := queryInteger(r, argumentConstant.Project)
	pipeline := queryInteger(r, argumentConstant.Pipeline)
	errors.PanicOnError(s.service.DeletePipeline(project, pipeline))
	s.worker.Poll()
	w.Header().Set(webConstant.ExtendedRedirect, constant.BoardPath)
}
