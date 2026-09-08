package model_context

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	mattermostFace "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	s *store.Store,
	i mattermostFace.Indexer,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:   m,
		monitor:  o,
		store:    s,
		indexer:  i,
		reporter: r,
	}
	result.register()

	return result
}
