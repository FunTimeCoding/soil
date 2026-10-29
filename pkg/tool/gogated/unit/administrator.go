package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func administrator(
	t *testing.T,
	b *base.Server,
	subject string,
) bool {
	t.Helper()
	result, e := b.Service.IsAdministrator(subject)
	assert.FatalOnError(t, e)

	return result
}
