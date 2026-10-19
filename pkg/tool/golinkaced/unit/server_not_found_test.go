package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/missing_link"
	"testing"
)

func TestAnUnknownLinkIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	s := server.New(missing_link.New(), nil, m)
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
