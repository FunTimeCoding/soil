package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"

func (o *Tester) Store() *store.Store {
	return o.server.Store
}
