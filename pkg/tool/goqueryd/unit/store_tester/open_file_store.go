package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func OpenFileStore(
	t *testing.T,
	path string,
) *store.Store {
	t.Helper()
	d := connection.New(logger.New(t.Context()), path)
	t.Cleanup(func() { errors.PanicClose(d) })

	return store.New(d, generativeConstant.OllamaEmbedModel)
}
