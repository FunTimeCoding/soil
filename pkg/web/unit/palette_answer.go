package unit

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"github.com/funtimecoding/soil/pkg/web/view"
	"net/http"
	"net/http/httptest"
	"testing"
)

func paletteAnswer(
	t *testing.T,
	v *view.View,
	bearer bool,
) *httptest.ResponseRecorder {
	t.Helper()
	m := http.NewServeMux()
	g := guard.New(m, []string{"token"}).WithSession(
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, q *http.Request) {
				next(w, guard.WithSubject(q, "alfa"))
			}
		},
	)
	g.Session(route.Get(constant.PalettePath), v.Palette)
	q := httptest.NewRequest(http.MethodGet, constant.PalettePath, nil)

	if bearer {
		web.Bearer(q, "token")
	}

	r := httptest.NewRecorder()
	m.ServeHTTP(r, q)

	return r
}
