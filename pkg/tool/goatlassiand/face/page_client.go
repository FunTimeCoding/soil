package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

type PageClient interface {
	ListSpaces(x context.Context) ([]*client.ConfluenceSpace, error)
	SearchPages(
		x context.Context,
		query string,
	) ([]*client.ConfluencePage, error)
	CreatePage(
		x context.Context,
		spaceIdentifier string,
		parentIdentifier string,
		title string,
		body string,
	) (*client.ConfluencePageDetail, error)
	GetPage(
		x context.Context,
		identifier string,
	) (*client.ConfluencePageDetail, error)
	UpdatePage(
		x context.Context,
		identifier string,
		title string,
		body string,
		message *string,
	) (*client.ConfluencePageDetail, error)
	DeletePage(
		x context.Context,
		identifier string,
		draft bool,
	) error
	GetPageChildren(
		x context.Context,
		identifier string,
	) ([]*client.ConfluencePage, error)
	AddPageComment(
		x context.Context,
		identifier string,
		body string,
	) error
	GetPageDraft(
		x context.Context,
		identifier string,
	) (*client.ConfluencePageDetail, error)
	SetPageStatus(
		x context.Context,
		identifier string,
		status string,
	) (*client.ConfluencePageDetail, error)
	EditPage(
		x context.Context,
		identifier string,
		oldText string,
		newText string,
		title string,
		message string,
		draft bool,
	) (*client.ConfluencePageDetail, error)
	ListPages(
		x context.Context,
		space string,
		status string,
	) ([]*client.ConfluencePage, error)
}
