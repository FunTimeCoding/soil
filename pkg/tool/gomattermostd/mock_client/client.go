package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

type Client struct {
	me     *model.User
	post   map[string]*model.Post
	thread map[string][]*post.Post
	user   map[string]*model.User
}
