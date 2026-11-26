package browser_tester

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"net/http"
	"net/http/httptest"
	"time"
)

func (b *Tester) Site() string {
	b.t.Helper()

	if b.site == nil {
		b.site = httptest.NewServer(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					switch q.URL.Path {
					case "/next":
						writer.Print(w, "%s", constant.FixtureNextPage)
					case "/slow":
						time.Sleep(500 * time.Millisecond)
						writer.Print(w, "%s", constant.FixtureNextPage)
					default:
						writer.Print(w, "%s", constant.FixtureLinkPage)
					}
				},
			),
		)
		b.t.Cleanup(b.site.Close)
	}

	return b.site.URL
}
