package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type MattermostSource interface {
	AllChannels(
		limit int,
		offset int,
	) ([]*model.ChannelWithTeamData, error)
	MustAllChannels(
		limit int,
		offset int,
	) []*model.ChannelWithTeamData
	Channel(identifier string) (*model.Channel, error)
	Channels(
		t *model.Team,
		u *model.User,
	) ([]*model.Channel, error)
	Context() context.Context
	DefaultTeam() *model.Team
	DirectMessage(
		u *model.User,
		text string,
	) (*model.Post, error)
	Enrich(v []*post.Post) error
	FindPost(identifier string) (*model.Post, error)
	LatestPosts(
		h *model.Channel,
		limit int,
	) ([]*post.Post, error)
	Me() (*model.User, error)
	MustMe() *model.User
	Nested() *model.Client4
	Post(p *model.Post) (*model.Post, error)
	PostSimple(
		h *model.Channel,
		text string,
	) (*model.Post, error)
	MustPostSimple(
		h *model.Channel,
		text string,
	) *model.Post
	PostsSince(
		h *model.Channel,
		since time.Time,
	) ([]*post.Post, error)
	MustPostsSince(
		h *model.Channel,
		since time.Time,
	) []*post.Post
	React(
		p *model.Post,
		emoji string,
	) error
	Reactions(p *model.Post) ([]*model.Reaction, error)
	Reply(
		h *model.Channel,
		m *model.Post,
		text string,
	) (*model.Post, error)
	TeamChannel(name string) (*model.Channel, error)
	MustTeamChannel(name string) *model.Channel
	Thread(p *model.Post) ([]*post.Post, error)
	User(identifier string) (*model.User, error)
	Users(
		page int,
		perPage int,
	) ([]*model.User, error)
}
