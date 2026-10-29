package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

type Client interface {
	Address() string
	Close()
	CreateLink(
		l string,
		title string,
		listIdentifier int,
		tags []string,
	) (*link.Link, error)
	CreateList(
		name string,
		description string,
	) (*list.List, error)
	CreateNote(
		linkIdentifier int,
		text string,
	) (*note.Note, error)
	CreateTag(name string) (*tag.Tag, error)
	DeleteLink(identifier int) error
	DeleteList(identifier int) error
	DeleteNote(identifier int) error
	DeleteTag(identifier int) error
	LinkByIdentifier(identifier int) (*link.Link, error)
	Links() ([]*link.Link, error)
	LinksByList(identifier int) ([]*link.Link, error)
	LinksByListPage(
		identifier int,
		p int,
	) (*page.Page[*link.Link], error)
	LinksPage(p int) (*page.Page[*link.Link], error)
	ListByName(name string) (*list.List, error)
	Lists() ([]*list.List, error)
	ListsPage(p int) (*page.Page[*list.List], error)
	NotesByLink(linkIdentifier int) ([]*note.Note, error)
	NotesByLinkPage(
		linkIdentifier int,
		p int,
	) (*page.Page[*note.Note], error)
	Reachable(x context.Context) error
	Search(query string) ([]*link.Link, error)
	SearchPage(
		query string,
		p int,
	) (*page.Page[*link.Link], error)
	TagByName(name string) (*tag.Tag, error)
	Tags() ([]*tag.Tag, error)
	TagsPage(p int) (*page.Page[*tag.Tag], error)
	UpdateLink(
		identifier int,
		patch map[string]any,
	) (*link.Link, error)
	UpdateList(
		identifier int,
		patch map[string]any,
	) (*list.List, error)
	UpdateNote(
		identifier int,
		text string,
	) (*note.Note, error)
	UpdateTag(
		identifier int,
		patch map[string]any,
	) (*tag.Tag, error)
}
