package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestExtendedRequestExcludesHistoryRestore(t *testing.T) {
	v := recoveryView()
	q := web.NewGet("/alfa")
	assert.False(t, v.IsExtendedRequest(q))
	q.Header.Set(constant.ExtendedRequest, "true")
	assert.True(t, v.IsExtendedRequest(q))
	q.Header.Set(constant.ExtendedRestore, "true")
	assert.False(t, v.IsExtendedRequest(q))
}
