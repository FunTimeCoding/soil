package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (o *Tester) OpenTestStore() *store.Store {
	o.t.Helper()
	d := connection.NewMemory()
	o.t.Cleanup(func() { errors.PanicClose(d) })

	return store.New(d, generativeConstant.OllamaEmbedModel)
}
