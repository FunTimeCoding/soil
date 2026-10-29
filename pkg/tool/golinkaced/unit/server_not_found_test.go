package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"testing"
)

func TestAnUnknownLinkIsAnsweredNotFound(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = not_found.Format("resource not found: links/%d", 12)
	m := memory.New()
	s := server.New(service_tester.Service(t, c, m), m)
	r, e := s.DeleteLink(
		context.Background(),
		generated.DeleteLinkRequestObject{Identifier: 12},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.DeleteLink404JSONResponse{
			Error: "resource not found: links/12",
		},
		r,
	)
	assert.Count(t, 0, m.Events())
}
