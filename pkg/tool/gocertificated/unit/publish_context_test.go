package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/publish_tester"
	"testing"
)

func TestPublishWhoseCallerLeftCommitsNothing(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	x, cancel := context.WithCancel(t.Context())
	cancel()
	_, _, e := o.Service.Publish(x)
	assert.True(t, errors.Is(e, context.Canceled))
	assert.Count(t, 0, o.Forge.Commits())
}
