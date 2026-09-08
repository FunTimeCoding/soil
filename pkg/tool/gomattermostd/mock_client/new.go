package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(me string) *Client {
	return &Client{
		me:     &model.User{Id: me, Username: me},
		post:   map[string]*model.Post{},
		thread: map[string][]*post.Post{},
		user:   map[string]*model.User{},
	}
}
