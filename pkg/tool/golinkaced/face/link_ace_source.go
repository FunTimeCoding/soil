package face

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

type LinkAceSource interface {
	Lists() ([]*list.List, error)
	ListsPage(p int) (*page.Page[*list.List], error)
	ListByName(name string) (*list.List, error)
	CreateList(
		name string,
		description string,
	) (*list.List, error)
	UpdateList(
		identifier int,
		patch map[string]any,
	) (*list.List, error)
	DeleteList(identifier int) error
	LinksPage(p int) (*page.Page[*link.Link], error)
	LinksByList(identifier int) ([]*link.Link, error)
	LinksByListPage(
		identifier int,
		p int,
	) (*page.Page[*link.Link], error)
	LinkByIdentifier(identifier int) (*link.Link, error)
	CreateLink(
		l string,
		title string,
		listIdentifier int,
		tags []string,
	) (*link.Link, error)
	UpdateLink(
		identifier int,
		patch map[string]any,
	) (*link.Link, error)
	DeleteLink(identifier int) error
	Search(query string) ([]*link.Link, error)
	Tags() ([]*tag.Tag, error)
	TagsPage(p int) (*page.Page[*tag.Tag], error)
	TagByName(name string) (*tag.Tag, error)
	CreateTag(name string) (*tag.Tag, error)
	UpdateTag(
		identifier int,
		patch map[string]any,
	) (*tag.Tag, error)
	DeleteTag(identifier int) error
	NotesByLinkPage(
		linkIdentifier int,
		p int,
	) (*page.Page[*note.Note], error)
	CreateNote(
		linkIdentifier int,
		text string,
	) (*note.Note, error)
	UpdateNote(
		identifier int,
		text string,
	) (*note.Note, error)
	DeleteNote(identifier int) error
}
