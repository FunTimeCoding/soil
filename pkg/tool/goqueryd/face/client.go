package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

type Client interface {
	GetSearch(
		x context.Context,
		p *client.GetSearchParams,
	) (*client.SearchOutcome, error)
	GetList(
		x context.Context,
		p *client.GetListParams,
	) (*client.ListOutcome, error)
	GetStatus(x context.Context) (*client.Status, error)
	GetOversize(
		x context.Context,
		p *client.GetOversizeParams,
	) (*client.OversizeReport, error)
	GetFailed(
		x context.Context,
		p *client.GetFailedParams,
	) (*client.FailedReport, error)
	PostChunk(
		x context.Context,
		b client.PostChunkJSONRequestBody,
	) (*client.ChunkPreview, error)
	PostRechunk(x context.Context) (*client.RechunkResult, error)
	PostIndex(
		x context.Context,
		b client.PostIndexJSONRequestBody,
	) ([]*client.IndexResult, error)
	PostEmbed(x context.Context) (*client.EmbedResult, error)
	GetContext(x context.Context) ([]*client.ContextEntry, error)
	GetDocument(
		x context.Context,
		p *client.GetDocumentParams,
	) (*client.SearchResult, error)
	PostDocument(
		x context.Context,
		b client.PostDocumentJSONRequestBody,
	) (*client.PushOutcome, error)
	GetMetadata(
		x context.Context,
		p *client.GetMetadataParams,
	) ([]*client.Facet, error)
	GetTags(x context.Context) ([]*client.SourceTypeTag, error)
	GetTag(
		x context.Context,
		p *client.GetTagParams,
	) (*client.SourceTypeTag, error)
	PostCollection(
		x context.Context,
		b client.PostCollectionJSONRequestBody,
	) error
	PostContext(
		x context.Context,
		b client.PostContextJSONRequestBody,
	) error
	PostTag(
		x context.Context,
		b client.PostTagJSONRequestBody,
	) error
	DeleteCollection(
		x context.Context,
		p *client.DeleteCollectionParams,
	) (bool, error)
	DeleteContext(
		x context.Context,
		p *client.DeleteContextParams,
	) (bool, error)
	DeleteDocument(
		x context.Context,
		p *client.DeleteDocumentParams,
	) (bool, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
