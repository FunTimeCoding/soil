package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"testing"
)

func holdStore(
	t *testing.T,
	path string,
) {
	t.Helper()
	holder := connection.New(logger.New(t.Context()), path)
	t.Cleanup(func() { errors.PanicClose(holder) })
	held, e := holder.BeginTx(t.Context(), nil)
	assert.FatalOnError(t, e)
	t.Cleanup(func() { assert.FatalOnError(t, held.Rollback()) })
}
