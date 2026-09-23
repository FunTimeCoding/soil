package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func corpus() *client.Statistics {
	return &client.Statistics{Scopes: []client.NamedCount{{Count: 500}}}
}
