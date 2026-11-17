package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/base"
	"testing"
)

func TestUnknownLinkIsAnsweredNotFound(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = not_found.Format("resource not found: links/%d", 12)
	r, e := s.GeneratedClient().DeleteLinkWithResponse(context.Background(), 12)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "resource not found: links/12", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}
