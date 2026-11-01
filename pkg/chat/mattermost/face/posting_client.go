package face

import "github.com/mattermost/mattermost/server/public/model"

type PostingClient interface {
	DeletePost(p *model.Post) error
	EditPost(
		p *model.Post,
		message string,
	) (*model.Post, error)
	MustDeletePost(p *model.Post)
	MustEditPost(
		p *model.Post,
		message string,
	) *model.Post
	MustPost(p *model.Post) *model.Post
	MustPostDefault(text string) *model.Post
	MustPostSimple(
		h *model.Channel,
		text string,
	) *model.Post
	MustReact(
		p *model.Post,
		emoji string,
	)
	MustReply(
		h *model.Channel,
		m *model.Post,
		text string,
	) *model.Post
	Post(p *model.Post) (*model.Post, error)
	PostDefault(text string) (*model.Post, error)
	PostSimple(
		h *model.Channel,
		text string,
	) (*model.Post, error)
	React(
		p *model.Post,
		emoji string,
	) error
	Reply(
		h *model.Channel,
		m *model.Post,
		text string,
	) (*model.Post, error)
	WipePrefix(
		h *model.Channel,
		prefix string,
	)
	WipePrefixKeepLatest(
		h *model.Channel,
		prefix string,
	)
}
