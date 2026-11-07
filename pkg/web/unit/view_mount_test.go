package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
	"net/http"
	"testing"
)

func TestTheViewMountsThePaletteBehindSignOnOnlyWhereThereIsSignOn(t *testing.T) {
	v := view.New(
		layout.New(identity.New("test", "test tool", "test"), "", []byte("png")).
			WithItems(navigation_item.New("/alfa", "Alfa")),
	)
	plain := http.NewServeMux()
	v.Mount(guard.New(plain, []string{"token"}))
	signed := http.NewServeMux()
	v.Mount(
		guard.New(signed, []string{"token"}).WithSession(
			func(http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(http.StatusForbidden)
				}
			},
		),
	)
	assert.Integer(t, http.StatusOK, served(plain, constant.PalettePath).Code)
	assert.Integer(
		t,
		http.StatusForbidden,
		served(signed, constant.PalettePath).Code,
	)
	assert.Integer(t, http.StatusOK, served(plain, constant.FaviconPath).Code)
	assert.Integer(t, http.StatusOK, served(signed, constant.FaviconPath).Code)
}
