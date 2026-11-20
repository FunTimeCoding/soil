package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
)

func (o *Tester) Open() (*store.Store, *ollama.Client) {
	o.t.Helper()
	c := o.client()

	return store_tester.New(o.t).OpenTestStore(), c
}
