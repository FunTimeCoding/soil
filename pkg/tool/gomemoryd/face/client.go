package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
	"time"
)

type Client interface {
	GetVersions(
		x context.Context,
		since string,
		limit int,
	) ([]version_entry.Entry, error)
	GetProfile(
		x context.Context,
		p *client.GetProfileParams,
	) (*client.ProfileResponse, error)
	PostMemory(
		x context.Context,
		b client.PostMemoryJSONRequestBody,
	) (*client.MemoryResponse, error)
	PutMemory(
		x context.Context,
		identifier int64,
		b client.PutMemoryJSONRequestBody,
	) (*client.MemoryResponse, error)
	DeleteMemory(
		x context.Context,
		identifier int64,
		p *client.DeleteMemoryParams,
	) (*client.MemoryResponse, error)
	GetSourcedMemories(
		x context.Context,
		p *client.GetSourcedMemoriesParams,
	) ([]client.SourcedMemory, error)
	GetRedactedMemories(x context.Context) (map[int64]bool, error)
	GetTokens(
		x context.Context,
		p *client.GetTokensParams,
	) (*client.TokenSummary, error)
	GetStatistics(x context.Context) (*client.Statistics, error)
	PostRelation(
		x context.Context,
		b client.PostRelationJSONRequestBody,
	) error
	DeleteRelation(
		x context.Context,
		p *client.DeleteRelationParams,
	) error
	GetRelations(
		x context.Context,
		p *client.GetRelationsParams,
	) ([]client.Relation, error)
	PostImpressions(
		x context.Context,
		content string,
		source string,
		capturedAt time.Time,
	) error
	Reachable(x context.Context) error
	Address() string
	Close()
}
