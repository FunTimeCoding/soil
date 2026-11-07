package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"github.com/funtimecoding/soil/pkg/web/view"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAFaviconIsLinkedFromTheHead(t *testing.T) {
	i := identity.New("test", "test tool", "test")
	assert.StringContains(
		t,
		`<link rel="icon" type="image/png" href="/favicon.ico">`,
		web_tester.RenderBrand(layout.New(i, "", []byte("png"))),
	)
	assert.StringNotContains(
		t,
		`rel="icon"`,
		web_tester.RenderBrand(layout.New(i, "", nil)),
	)
}

func TestTheViewServesTheFaviconItHolds(t *testing.T) {
	i := identity.New("test", "test tool", "test")
	r := httptest.NewRecorder()
	view.New(layout.New(i, "", []byte("png"))).Favicon(
		r,
		httptest.NewRequest(http.MethodGet, constant.FaviconPath, nil),
	)
	assert.Integer(t, http.StatusOK, r.Code)
	assert.String(t, "image/png", r.Header().Get(constant.ContentType))
	assert.String(t, "png", r.Body.String())
	none := httptest.NewRecorder()
	view.New(layout.New(i, "", nil)).Favicon(
		none,
		httptest.NewRequest(http.MethodGet, constant.FaviconPath, nil),
	)
	assert.Integer(t, http.StatusNotFound, none.Code)
}
