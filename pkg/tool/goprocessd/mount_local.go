package goprocessd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func MountLocal(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
	m *http.ServeMux,
) {
	m.HandleFunc(
		route.Get(constant.ReachablePath),
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.WriteHeader(http.StatusNoContent)
		},
	)
	m.Handle(constant.InterfacePath, interfaceTree(v, r, t))
}
