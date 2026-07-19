package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/request"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestJobToken(t *testing.T) {
	r := web.NewGet(locator.New(constant.Localhost).Insecure().String())
	request.JobToken(r, "test")
	assert.Any(t, http.Header{"Job-Token": {"test"}}, r.Header)
}
