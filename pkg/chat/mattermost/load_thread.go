package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
)

func (c *Client) LoadThread(p *post.Post) *thread.Thread {
	result := thread.New(p, c.MustThread(p.Raw))
	result.Resolved = c.MustPostHasReaction(p.Raw, constant.CheckMark)

	return result
}
