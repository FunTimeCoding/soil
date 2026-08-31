package web

import (
	argument "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	project := queryInteger(r, argument.Project)
	pipeline := queryInteger(r, argument.Pipeline)
	errors.PanicOnError(s.client.DeletePipeline(project, pipeline))
	s.worker.Poll()
	w.Header().Set(web.ExtendedRedirect, constant.BoardPath)
}
