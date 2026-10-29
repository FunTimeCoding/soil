package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func OpenTestStore(t *testing.T) *store.Store {
	t.Helper()
	d := connection.NewMemory()
	t.Cleanup(func() { errors.PanicClose(d) })

	return store.New(d, generativeConstant.OllamaEmbedModel)
}
