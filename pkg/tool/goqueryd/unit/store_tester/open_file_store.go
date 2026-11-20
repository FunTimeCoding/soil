package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (o *Tester) OpenFileStore(path string) *store.Store {
	o.t.Helper()
	d := connection.New(logger.New(o.t.Context()), path)
	o.t.Cleanup(func() { errors.PanicClose(d) })

	return store.New(d, generativeConstant.OllamaEmbedModel)
}
