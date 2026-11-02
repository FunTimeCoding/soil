package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type SnippetClient interface {
	CreateSnippet(
		x context.Context,
		p *client.CreateSnippetParams,
		b client.CreateSnippetJSONRequestBody,
	) (*client.SnippetCreateResult, error)
	DeleteSnippet(
		x context.Context,
		name string,
		p *client.DeleteSnippetParams,
	) (*client.StatusResult, error)
	GetSnippet(
		x context.Context,
		name string,
		p *client.GetSnippetParams,
	) (*client.SnippetDetail, error)
	ListSnippets(
		x context.Context,
		p *client.ListSnippetsParams,
	) ([]client.SnippetItem, error)
}
