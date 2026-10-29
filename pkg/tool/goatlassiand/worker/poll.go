package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"slices"
)

func (w *Worker) Poll() {
	p, e := w.service.Plate()

	if e != nil {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)
		} else {
			w.logger.Plain(constant.PlateDeferred, e)
		}

		return
	}

	w.mutex.Lock()
	previous := w.issues
	previousWatchedIssues := w.watchedIssues
	previousNewest := w.newest
	previousFavorites := w.favorites
	previousWatched := w.watched
	w.issues = p.Issues
	w.watchedIssues = p.WatchedIssues
	w.newest = p.Newest
	w.favorites = p.Favorites
	w.watched = p.Watched
	w.mutex.Unlock()

	if !slices.EqualFunc(previous, p.Issues, sameIssue) ||
		!slices.EqualFunc(previousWatchedIssues, p.WatchedIssues, sameIssue) ||
		!slices.EqualFunc(previousNewest, p.Newest, sameIssue) ||
		!slices.EqualFunc(previousFavorites, p.Favorites, samePage) ||
		!slices.EqualFunc(previousWatched, p.Watched, samePage) {
		w.notifier.Notify()
	}
}
