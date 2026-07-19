package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestBearer(t *testing.T) {
	r := web.NewGet(constant.Localhost)
	web.Bearer(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Bearer Alfa"}},
		r.Header,
	)
}
