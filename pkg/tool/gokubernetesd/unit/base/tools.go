package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
)

func (s *Server) Tools() *model_context_client.Client {
	s.t.Helper()
	result := s.Served().Connect(s.t)
	result.MustCallTool(
		constant.UseCluster,
		map[string]any{"cluster": s.Cluster.Name()},
	)

	return result
}
