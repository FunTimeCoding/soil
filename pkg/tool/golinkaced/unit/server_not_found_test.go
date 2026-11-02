package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"testing"
)

func TestAnUnknownLinkIsAnsweredNotFound(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = not_found.Format("resource not found: links/%d", 12)
	m := memory.New()
	r, e := restClient(t, c, m).DeleteLinkWithResponse(context.Background(), 12)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "resource not found: links/12", r.JSON404.Error)
	assert.Count(t, 0, m.Events())
}
