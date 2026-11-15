package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/unit/view_tester"
	"testing"
)

func TestExtendedRequestExcludesHistoryRestore(t *testing.T) {
	v := view_tester.New()
	q := web.NewGet("/alfa")
	assert.False(t, v.IsExtendedRequest(q))
	q.Header.Set(constant.ExtendedRequest, "true")
	assert.True(t, v.IsExtendedRequest(q))
	q.Header.Set(constant.ExtendedRestore, "true")
	assert.False(t, v.IsExtendedRequest(q))
}
