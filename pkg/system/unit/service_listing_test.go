package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestFailedListingIsErrorNotEmptyHost(t *testing.T) {
	t.Setenv(constant.PathEnvironment, t.TempDir())
	v, e := service.New().Services()
	assert.Error(t, e)
	assert.Count(t, 0, v)
}
