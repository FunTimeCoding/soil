package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestToken(t *testing.T) {
	r := NewGet(constant.Localhost)
	Token(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Token Alfa"}},
		r.Header,
	)
}
