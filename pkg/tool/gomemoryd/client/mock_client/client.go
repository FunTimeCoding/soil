package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

type Client struct {
	Impressions []ImpressionCall
	Redacted    map[int64]bool
	Stats       *client.Statistics
	Edges       []client.Relation
}
