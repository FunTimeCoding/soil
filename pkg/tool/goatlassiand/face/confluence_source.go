package face

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
)

type ConfluenceSource interface {
	Page(identifier string) (*page.Page, error)
	DraftPage(identifier string) (*page.Page, error)
	DraftOverlay(identifier string) (*page.Page, error)
	Search(
		query string,
		a ...any,
	) ([]*search_result.Result, error)
	Spaces() ([]*space.Space, error)
	ChildPagesByIdentifier(identifier string) ([]*page.Page, error)
	CreatePage(
		spaceIdentifier string,
		parentIdentifier string,
		title string,
		markdown string,
	) (*page.Page, error)
	CreateDraftPage(
		spaceIdentifier string,
		parentIdentifier string,
		title string,
		markdown string,
	) (*page.Page, error)
	PutPage(
		identifier string,
		title string,
		body string,
		version int,
		message string,
		status string,
	) (*page.Page, error)
	DraftPages() ([]*page.Page, error)
	PagesBySpace(
		identifier string,
		status string,
	) ([]*page.Page, error)
	AddComment(
		pageIdentifier string,
		body string,
	) error
	Delete(pageIdentifier string) error
	DeleteDraft(pageIdentifier string) error
}
