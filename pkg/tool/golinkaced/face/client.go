package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

type Client interface {
	AddLink(
		x context.Context,
		b client.AddLinkJSONRequestBody,
	) (*client.Link, error)
	AddNote(
		x context.Context,
		b client.AddNoteJSONRequestBody,
	) (*client.Note, error)
	AppendList(
		x context.Context,
		identifier int32,
		b client.AppendListJSONRequestBody,
	) (*client.Link, error)
	AppendTag(
		x context.Context,
		identifier int32,
		b client.AppendTagJSONRequestBody,
	) (*client.Link, error)
	CreateList(
		x context.Context,
		b client.CreateListJSONRequestBody,
	) (*client.List, error)
	CreateTag(
		x context.Context,
		b client.CreateTagJSONRequestBody,
	) (*client.Tag, error)
	DeleteBranch(
		x context.Context,
		identifier int32,
	) error
	DeleteLink(
		x context.Context,
		identifier int32,
	) error
	DeleteList(
		x context.Context,
		identifier int32,
	) error
	DeleteNote(
		x context.Context,
		identifier int32,
	) error
	DeleteTag(
		x context.Context,
		identifier int32,
	) error
	EditLink(
		x context.Context,
		identifier int32,
		b client.EditLinkJSONRequestBody,
	) (*client.Link, error)
	EditList(
		x context.Context,
		identifier int32,
		b client.EditListJSONRequestBody,
	) (*client.List, error)
	EditNote(
		x context.Context,
		identifier int32,
		b client.EditNoteJSONRequestBody,
	) (*client.Note, error)
	EditTag(
		x context.Context,
		identifier int32,
		b client.EditTagJSONRequestBody,
	) (*client.Tag, error)
	GetLinkNotes(
		x context.Context,
		identifier int32,
		p *client.GetLinkNotesParams,
	) (*client.NotePage, error)
	GetLinks(
		x context.Context,
		p *client.GetLinksParams,
	) (*client.LinkPage, error)
	GetLists(
		x context.Context,
		p *client.GetListsParams,
	) (*client.ListPage, error)
	GetTags(
		x context.Context,
		p *client.GetTagsParams,
	) (*client.TagPage, error)
	RemoveList(
		x context.Context,
		identifier int32,
		b client.RemoveListJSONRequestBody,
	) (*client.Link, error)
	RemoveTag(
		x context.Context,
		identifier int32,
		b client.RemoveTagJSONRequestBody,
	) (*client.Link, error)
	Search(
		x context.Context,
		p *client.SearchParams,
	) (*client.SearchResult, error)
	Reachable(x context.Context) error
	Close()
}
