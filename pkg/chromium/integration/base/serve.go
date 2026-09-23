package base

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"net/http"
)

func serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	switch q.URL.Path {
	case constant.FixtureBusyRoute:
		write(w, constant.FixtureBusyPage)
	case constant.FixtureStalledRoute:
		write(w, constant.FixtureStalledPage)
	case constant.FixtureUnloadRoute:
		write(w, constant.FixtureUnloadPage)
	case constant.FixturePopupRoute:
		write(w, constant.FixturePopupPage)
	case constant.FixtureHeavyRoute:
		write(w, heavyPage())
	case constant.FixturePingRoute:
		w.WriteHeader(http.StatusNoContent)
	case constant.FixtureHangRoute:
		<-q.Context().Done()
	default:
		write(w, constant.FixtureQuietPage)
	}
}
