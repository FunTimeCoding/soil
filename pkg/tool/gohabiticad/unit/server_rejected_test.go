package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestInvalidStatIsAnsweredAsBadRequest(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("%q is not a valid Stat.", "bravo")
	m := memory.New()
	r, e := restClient(t, c, m).AllocateStatWithResponse(
		context.Background(),
		client.AllocateStatParamsStat("bravo"),
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, `"bravo" is not a valid Stat.`, r.JSON400.Error)
	assert.Count(t, 0, m.Events())
}
