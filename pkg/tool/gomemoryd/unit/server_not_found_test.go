package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestAnUnknownMemoryIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	r, e := server.New(service_tester.New(t).Service, m).DeleteMemory(
		context.Background(),
		generated.DeleteMemoryRequestObject{Identifier: 999},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.DeleteMemory404JSONResponse{Error: "memory not found: 999"},
		r,
	)
	assert.Count(t, 0, m.Events())
}
