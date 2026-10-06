package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func sameResult(
	t *testing.T,
	run func(s *service.Service) (any, any, error),
	indexed *service.Service,
	full *service.Service,
) string {
	t.Helper()
	r1, v1, e1 := run(indexed)
	r2, v2, e2 := run(full)
	assert.FatalOnError(t, e1)
	assert.FatalOnError(t, e2)
	result := string(notation.Marshal([]any{r1, v1}))
	assert.String(t, string(notation.Marshal([]any{r2, v2})), result)

	return result
}
