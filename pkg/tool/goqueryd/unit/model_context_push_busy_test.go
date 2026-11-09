package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestPushToolWhileAnotherWriterHoldsStoreFailsReported(t *testing.T) {
	r := memory.New()
	o := service_tester.NewOption(t, newEmbedder(t), newReranker(t))
	c := served(t, o, r).Connect(t)
	defer c.Close()
	holdStore(t, o.LitePath)
	assert.StringContains(
		t,
		"push failed",
		c.MustCallToolError(
			constant.Push,
			map[string]any{
				constant.Collection: "notes",
				constant.Path:       "alfa.md",
				constant.Body:       "# Alfa\n\nzulu.\n",
			},
		),
	)
	assert.Count(t, 1, r.Events())
}
