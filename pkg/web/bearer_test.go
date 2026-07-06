package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestBearer(t *testing.T) {
	r := NewGet(constant.Localhost)
	Bearer(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Bearer Alfa"}},
		r.Header,
	)
}
