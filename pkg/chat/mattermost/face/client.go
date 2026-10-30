package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type Client interface {
	AllChannels(
		limit int,
		offset int,
	) ([]*model.ChannelWithTeamData, error)
	Channel(identifier string) (*model.Channel, error)
	ChannelByName(
		t *model.Team,
		name string,
	) (*model.Channel, error)
	ChannelUsers(h *model.Channel) ([]*model.User, error)
	Channels(
		t *model.Team,
		u *model.User,
	) ([]*model.Channel, error)
	Close()
	Configuration() (map[string]string, error)
	DefaultChannel() *model.Channel
	DefaultTeam() *model.Team
	DeletePost(p *model.Post) error
	EditPost(
		p *model.Post,
		message string,
	) (*model.Post, error)
	DirectMessage(
		u *model.User,
		text string,
	) (*model.Post, error)
	Enrich(v []*post.Post) error
	File(identifier string) ([]byte, error)
	FileInformation(identifier string) (*model.FileInfo, error)
	FindPost(identifier string) (*model.Post, error)
	FindPostBefore(
		h *model.Channel,
		t time.Time,
	) (*post.Post, bool, error)
	Host() string
	LatestPosts(
		h *model.Channel,
		limit int,
	) ([]*post.Post, error)
	LoadThread(p *post.Post) *thread.Thread
	Me() (*model.User, error)
	MustAllChannels(
		limit int,
		offset int,
	) []*model.ChannelWithTeamData
	MustChannel(identifier string) *model.Channel
	MustChannelByName(
		t *model.Team,
		name string,
	) *model.Channel
	MustChannelUsers(h *model.Channel) []*model.User
	MustChannels(
		t *model.Team,
		u *model.User,
	) []*model.Channel
	MustConfiguration() map[string]string
	MustDeletePost(p *model.Post)
	MustEditPost(
		p *model.Post,
		message string,
	) *model.Post
	MustDirectMessage(
		u *model.User,
		text string,
	) *model.Post
	MustEnrich(v []*post.Post)
	MustFindPost(identifier string) *model.Post
	MustLatestPosts(
		h *model.Channel,
		limit int,
	) []*post.Post
	MustMe() *model.User
	MustOpenDialog(v model.OpenDialogRequest)
	MustPost(p *model.Post) *model.Post
	MustPostBefore(
		h *model.Channel,
		t time.Time,
	) *post.Post
	MustPostBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
	) *post.Post
	MustPostDefault(text string) *model.Post
	MustPostHasReaction(
		p *model.Post,
		reaction string,
	) bool
	MustPostSimple(
		h *model.Channel,
		text string,
	) *model.Post
	MustPosts(h *model.Channel) *model.PostList
	MustPostsBefore(
		h *model.Channel,
		t time.Time,
		keep int,
	) []*post.Post
	MustPostsBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
		keep int,
	) []*post.Post
	MustPostsSince(
		h *model.Channel,
		since time.Time,
	) []*post.Post
	MustReact(
		p *model.Post,
		emoji string,
	)
	MustReactions(p *model.Post) []*model.Reaction
	MustRefreshSocket()
	MustReply(
		h *model.Channel,
		m *model.Post,
		text string,
	) *model.Post
	MustSubmitDialog(v model.SubmitDialogRequest) *model.SubmitDialogResponse
	MustTeam(name string) *model.Team
	MustTeamChannel(name string) *model.Channel
	MustTeams(userIdentifier string) []*model.Team
	MustThread(p *model.Post) []*post.Post
	MustUpdateChannelHeader(
		h *model.Channel,
		prefix string,
		text string,
		separator string,
	) *model.Channel
	MustUser(identifier string) *model.User
	MustUsers(
		page int,
		perPage int,
	) []*model.User
	OpenDialog(v model.OpenDialogRequest) error
	Post(p *model.Post) (*model.Post, error)
	PostBefore(
		h *model.Channel,
		t time.Time,
	) (*post.Post, error)
	PostBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
	) (*post.Post, error)
	PostDefault(text string) (*model.Post, error)
	PostHasReaction(
		p *model.Post,
		reaction string,
	) (bool, error)
	PostSimple(
		h *model.Channel,
		text string,
	) (*model.Post, error)
	Posts(h *model.Channel) (*model.PostList, error)
	PostsBefore(
		h *model.Channel,
		t time.Time,
		keep int,
	) ([]*post.Post, error)
	PostsBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
		keep int,
	) ([]*post.Post, error)
	PostsSince(
		h *model.Channel,
		since time.Time,
	) ([]*post.Post, error)
	Reachable(x context.Context) error
	React(
		p *model.Post,
		emoji string,
	) error
	Reactions(p *model.Post) ([]*model.Reaction, error)
	RefreshSocket() error
	Reply(
		h *model.Channel,
		m *model.Post,
		text string,
	) (*model.Post, error)
	SearchChannels(
		t *model.Team,
		query string,
	) ([]*model.Channel, error)
	SearchPosts(
		t *model.Team,
		terms string,
	) (*model.PostList, error)
	SearchUsers(
		t *model.Team,
		query string,
		limit int,
	) ([]*model.User, error)
	SubmitDialog(v model.SubmitDialogRequest) (*model.SubmitDialogResponse, error)
	Team(name string) (*model.Team, error)
	TeamChannel(name string) (*model.Channel, error)
	Teams(userIdentifier string) ([]*model.Team, error)
	Thread(p *model.Post) ([]*post.Post, error)
	UpdateChannelHeader(
		h *model.Channel,
		prefix string,
		text string,
		separator string,
	) (*model.Channel, error)
	UploadFile(
		h *model.Channel,
		name string,
		content []byte,
	) (*model.FileInfo, error)
	User(identifier string) (*model.User, error)
	UserByName(name string) (*model.User, error)
	UserMap() *user_map.Map
	UserThreads(
		u *model.User,
		t *model.Team,
		o model.GetUserThreadsOpts,
	) (*model.Threads, error)
	Users(
		page int,
		perPage int,
	) ([]*model.User, error)
	WebSocket() *model.WebSocketClient
	WipePrefix(
		h *model.Channel,
		prefix string,
	)
	WipePrefixKeepLatest(
		h *model.Channel,
		prefix string,
	)
}
