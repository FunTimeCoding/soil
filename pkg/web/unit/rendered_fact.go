package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"strings"
	"testing"
)

func renderedFact(
	t *testing.T,
	f *fact.Fact,
) string {
	t.Helper()
	var b strings.Builder
	assert.FatalOnError(t, f.Node().Render(&b))

	return b.String()
}
