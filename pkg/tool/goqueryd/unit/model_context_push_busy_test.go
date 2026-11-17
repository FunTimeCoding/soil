package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/held_store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestPushToolWhileAnotherWriterHoldsStoreFailsReported(t *testing.T) {
	q := service_tester.New(t)
	c := q.Served().Connect(t)
	defer c.Close()
	held_store.New(t, q.Path)
	assert.StringContains(
		t,
		"store busy",
		c.MustCallToolError(
			constant.Push,
			map[string]any{
				constant.Collection: "notes",
				constant.Path:       "alfa.md",
				constant.Body:       "# Alfa\n\nzulu.\n",
			},
		),
	)
	assert.Count(t, 1, q.Reporter.Events())
}
