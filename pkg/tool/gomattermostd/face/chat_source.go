package face

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

type ChatSource interface {
	Me() (*model.User, error)
	User(identifier string) (*model.User, error)
	FindPost(identifier string) (*model.Post, error)
	Thread(p *model.Post) ([]*post.Post, error)
	WebSocket() *model.WebSocketClient
	RefreshSocket()
}
