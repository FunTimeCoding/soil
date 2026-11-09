package face

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type MattermostPostReadingClient interface {
	Enrich(v []*post.Post) error
	FindPost(identifier string) (*model.Post, error)
	LoadThread(p *post.Post) *thread.Thread
	MustEnrich(v []*post.Post)
	MustFindPost(identifier string) *model.Post
	MustPostHasReaction(
		p *model.Post,
		reaction string,
	) bool
	MustPosts(
		h *model.Channel,
		since time.Time,
	) []*post.Post
	MustReactions(p *model.Post) []*model.Reaction
	MustRecentPosts(
		h *model.Channel,
		limit int,
	) []*post.Post
	MustThread(p *model.Post) []*post.Post
	PostHasReaction(
		p *model.Post,
		reaction string,
	) (bool, error)
	Posts(
		h *model.Channel,
		since time.Time,
	) ([]*post.Post, error)
	Reactions(p *model.Post) ([]*model.Reaction, error)
	RecentPosts(
		h *model.Channel,
		limit int,
	) ([]*post.Post, error)
	SearchPosts(
		t *model.Team,
		terms string,
	) (*model.PostList, error)
	Thread(p *model.Post) ([]*post.Post, error)
	UserThreads(
		u *model.User,
		t *model.Team,
		o model.GetUserThreadsOpts,
	) (*model.Threads, error)
}
