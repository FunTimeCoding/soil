package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"testing"
)

func TestInvalidStatIsAnsweredAsBadRequest(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = validation.New("%q is not a valid Stat.", "bravo")
	r, e := s.GeneratedClient().AllocateStatWithResponse(
		context.Background(),
		client.AllocateStatParamsStat("bravo"),
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, `"bravo" is not a valid Stat.`, r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}
