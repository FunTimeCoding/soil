package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) SameResult(
	run func(s *service_tester.Services) (any, any, error),
	indexed *service_tester.Services,
	full *service_tester.Services,
) string {
	o.t.Helper()
	r1, v1, e1 := run(indexed)
	r2, v2, e2 := run(full)
	assert.FatalOnError(o.t, e1)
	assert.FatalOnError(o.t, e2)
	result := string(notation.Marshal([]any{r1, v1}))
	assert.String(o.t, string(notation.Marshal([]any{r2, v2})), result)

	return result
}
