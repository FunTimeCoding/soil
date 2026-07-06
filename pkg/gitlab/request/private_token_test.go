package request

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestPrivateToken(t *testing.T) {
	r := web.NewGet(locator.New(constant.Localhost).Insecure().String())
	PrivateToken(r, "test")
	assert.Any(t, http.Header{"Private-Token": {"test"}}, r.Header)
}
