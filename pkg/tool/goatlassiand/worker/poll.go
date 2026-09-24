package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"slices"
)

func (w *Worker) Poll() {
	closed := w.closedStatus()
	issues, e := w.client.Search(constant.PlateQuery, closed)
	errors.PanicOnError(e)
	watchedIssues, f := w.client.Search(constant.WatchedIssuesQuery, closed)
	errors.PanicOnError(f)
	newest := w.fetchNewest(closed)
	favorites, g := w.confluence.Favorites()
	errors.PanicOnError(g)
	watched, h := w.confluence.Watched()
	errors.PanicOnError(h)
	w.mutex.Lock()
	previous := w.issues
	previousWatchedIssues := w.watchedIssues
	previousNewest := w.newest
	previousFavorites := w.favorites
	previousWatched := w.watched
	w.issues = issues
	w.watchedIssues = watchedIssues
	w.newest = newest
	w.favorites = favorites
	w.watched = watched
	w.mutex.Unlock()

	if !slices.EqualFunc(previous, issues, sameIssue) ||
		!slices.EqualFunc(previousWatchedIssues, watchedIssues, sameIssue) ||
		!slices.EqualFunc(previousNewest, newest, sameIssue) ||
		!slices.EqualFunc(previousFavorites, favorites, samePage) ||
		!slices.EqualFunc(previousWatched, watched, samePage) {
		w.notifier.Notify()
	}
}
