package channel

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ConfirmChannel(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ConfirmChannel,
) (*mcp.CallToolResult, error) {
	if a.Nonce != s.nonce {
		return response.Fail(constant.ChannelNonceMismatch)
	}

	if a.Callsign == "" {
		return response.Fail(constant.ChannelCallsignRequired)
	}

	s.mutex.Lock()
	resolved := s.callsign

	if resolved != "" && resolved != a.Callsign {
		s.mutex.Unlock()

		return response.Fail(
			constant.ChannelCallsignMismatch,
			resolved,
			a.Callsign,
		)
	}

	s.callsign = a.Callsign
	s.mutex.Unlock()
	s.openOnce.Do(func() { close(s.open) })

	return response.Success(constant.ChannelOpened)
}
