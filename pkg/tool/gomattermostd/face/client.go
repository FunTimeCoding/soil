package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

type Client interface {
	GetStatus(x context.Context) (*client.Status, error)
	GetChannelPosts(
		x context.Context,
		name string,
		p *client.GetChannelPostsParams,
	) ([]*client.Post, error)
	GetUnresolvedThreads(
		x context.Context,
		name string,
		p *client.GetUnresolvedThreadsParams,
	) ([]client.Thread, error)
	GetPostReactions(
		x context.Context,
		identifier string,
	) ([]string, error)
	GetThread(
		x context.Context,
		identifier string,
	) ([]*client.Post, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
