package context_panel_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func Corpus() *client.Statistics {
	return &client.Statistics{Scopes: []client.NamedCount{{Count: 500}}}
}
