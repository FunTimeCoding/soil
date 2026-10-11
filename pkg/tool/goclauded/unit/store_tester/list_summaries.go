package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"

func (o *Tester) ListSummaries() []summary.Summary {
	return o.Store.ListSummaries()
}
