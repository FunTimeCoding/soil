package mock_client

import "github.com/funtimecoding/soil/pkg/generative/types/tool_context_result"

func (c *Client) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []tool_context_result.Result {
	return nil
}
