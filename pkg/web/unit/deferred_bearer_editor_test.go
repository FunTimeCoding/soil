package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestDeferredBearerEditorReadsTheTokenPerRequest(t *testing.T) {
	var names []string
	read := func(name string) string {
		names = append(names, name)

		return "bravo"
	}
	e := web.DeferredBearerEditor(read, "ALFA_TOKEN")
	assert.Count(t, 0, names)
	q, f := http.NewRequest(http.MethodGet, "http://localhost", nil)
	assert.FatalOnError(t, f)
	assert.FatalOnError(t, e(context.Background(), q))
	assert.String(t, "Bearer bravo", q.Header.Get(constant.Authorization))
	assert.Strings(t, []string{"ALFA_TOKEN"}, names)
}
