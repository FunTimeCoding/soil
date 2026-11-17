package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func New(t *testing.T, me string) *Client {
	t.Helper()
	result := &Client{
		me:     &model.User{Id: me, Username: me},
		post:   map[string]*model.Post{},
		thread: map[string][]*post.Post{},
		user:   map[string]*model.User{},
	}
	t.Cleanup(result.Close)

	return result
}
