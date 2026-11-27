package base

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"net/http"
)

func serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	switch q.URL.Path {
	case constant.ChromiumFixtureBusyRoute:
		write(w, constant.ChromiumFixtureBusyPage)
	case constant.ChromiumFixtureStalledRoute:
		write(w, constant.ChromiumFixtureStalledPage)
	case constant.ChromiumFixtureUnloadRoute:
		write(w, constant.ChromiumFixtureUnloadPage)
	case constant.ChromiumFixturePopupRoute:
		write(w, constant.ChromiumFixturePopupPage)
	case constant.ChromiumFixtureHeavyRoute:
		write(w, heavyPage())
	case constant.ChromiumFixturePingRoute:
		w.WriteHeader(http.StatusNoContent)
	case constant.ChromiumFixtureHangRoute:
		<-q.Context().Done()
	default:
		write(w, constant.ChromiumFixtureQuietPage)
	}
}
