package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/impression_call"
)

type Client struct {
	Impressions []impression_call.Call
	Redacted    map[int64]bool
	Stats       *client.Statistics
	Edges       []client.Relation
}
