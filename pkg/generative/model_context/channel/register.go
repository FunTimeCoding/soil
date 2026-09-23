package channel

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ChannelConfirmTool,
			mcp.WithDescription(constant.ChannelConfirmDescription),
			mcp.WithString(
				constant.ChannelNonceMeta,
				mcp.Description(constant.ChannelNonceDescription),
				mcp.Required(),
			),
			mcp.WithString(
				constant.ChannelCallsignMeta,
				mcp.Description(constant.ChannelCallsignDescription),
				mcp.Required(),
			),
		),
		mcp.NewTypedToolHandler(s.ConfirmChannel),
	)
}
