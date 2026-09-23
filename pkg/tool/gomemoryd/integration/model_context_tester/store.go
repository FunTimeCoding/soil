package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (o *Tester) Store() *store.Store {
	return o.base.Store()
}
