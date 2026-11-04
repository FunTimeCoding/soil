package model_context

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
)

func replyRow(r *post.Post) *response.Reply {
	result := response.NewReply(r.Raw.Id, r.Message, formatTime(r.Create))

	if r.User != nil {
		result.Username = r.User.Username
	}

	if len(r.Raw.FileIds) > 0 {
		result.FileIds = r.Raw.FileIds
	}

	return result
}
