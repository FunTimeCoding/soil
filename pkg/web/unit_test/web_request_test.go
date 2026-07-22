package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestAuthorization(t *testing.T) {
	r := web.NewGet(constant.Localhost)
	web.Bearer(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Bearer Alfa"}},
		r.Header,
	)
	r = web.NewGet(constant.Localhost)
	web.Token(r, upper.Alfa)
	assert.Any(
		t,
		http.Header{"Authorization": {"Token Alfa"}},
		r.Header,
	)
}

func TestGetList(t *testing.T) {
	assert.Any(
		t,
		[]string{"1", "2", "3"},
		web.GetList(
			web.NewGet(
				locator.New(
					constant.Localhost,
				).Insecure().Set("a", "1,2,3").String(),
			),
			"a",
		),
	)
}

func TestClient(t *testing.T) {
	assert.Any(t, &http.Client{}, web.Client())
}
