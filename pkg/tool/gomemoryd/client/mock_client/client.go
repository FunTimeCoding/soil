package mock_client

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/impression_call"
)

type Client struct {
	Impressions []impression_call.Call
	Retargets   []*move_set.Set
	Redacted    map[int64]bool
	Stats       *client.Statistics
	Edges       []client.Relation
	Failure     error
	Closed      int
}
