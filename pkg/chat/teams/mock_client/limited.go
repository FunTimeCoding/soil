package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
)

func limited(
	v []*message.Message,
	limit int,
) []*message.Message {
	if limit <= 0 {
		limit = constant.TeamsDefaultLimit
	}

	return append([]*message.Message{}, v[:min(limit, len(v))]...)
}
