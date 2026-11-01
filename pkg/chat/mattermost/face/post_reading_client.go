package face

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type PostReadingClient interface {
	Enrich(v []*post.Post) error
	FindPost(identifier string) (*model.Post, error)
	FindPostBefore(
		h *model.Channel,
		t time.Time,
	) (*post.Post, bool, error)
	LatestPosts(
		h *model.Channel,
		limit int,
	) ([]*post.Post, error)
	LoadThread(p *post.Post) *thread.Thread
	MustEnrich(v []*post.Post)
	MustFindPost(identifier string) *model.Post
	MustLatestPosts(
		h *model.Channel,
		limit int,
	) []*post.Post
	MustPostBefore(
		h *model.Channel,
		t time.Time,
	) *post.Post
	MustPostBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
	) *post.Post
	MustPostHasReaction(
		p *model.Post,
		reaction string,
	) bool
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
	MustReactions(p *model.Post) []*model.Reaction
	MustThread(p *model.Post) []*post.Post
	PostBefore(
		h *model.Channel,
		t time.Time,
	) (*post.Post, error)
	PostBeforeMilli(
		h *model.Channel,
		beforeMilli int64,
	) (*post.Post, error)
	PostHasReaction(
		p *model.Post,
		reaction string,
	) (bool, error)
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
	Reactions(p *model.Post) ([]*model.Reaction, error)
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
