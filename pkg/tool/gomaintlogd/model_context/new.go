package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
)

func New(
	s *store.Store,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity).WithRecorder(t).Server(),
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
