package channel_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func Confirm(
	s *channel.Server,
	nonce string,
	callsign string,
) *mcp.CallToolResult {
	result, e := s.ConfirmChannel(
		context.Background(),
		mcp.CallToolRequest{},
		argument.ConfirmChannel{Nonce: nonce, Callsign: callsign},
	)

	if e != nil {
		panic(e)
	}

	return result
}
