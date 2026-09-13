package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) fail(
	_ http.ResponseWriter,
	_ *http.Request,
) {
	panic(constant.FailMessage)
}
