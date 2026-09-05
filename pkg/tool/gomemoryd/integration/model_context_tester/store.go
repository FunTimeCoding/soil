package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (t *Tester) Store() *store.Store {
	return t.base.Store()
}
