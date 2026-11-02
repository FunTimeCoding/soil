package usage

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) WithUtilization(
	c face.AnthropicClient,
	g *peer.Guard,
	read func() *credential.Credential,
) *Service {
	s.utilization = c
	s.utilizationGuard = g
	s.credential = read

	return s
}
