package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"slices"
)

func (w *Worker) Poll() {
	issues, e := w.client.Search(constant.PlateQuery)
	errors.PanicOnError(e)
	favorites, f := w.confluence.Favorites()
	errors.PanicOnError(f)
	watched, g := w.confluence.Watched()
	errors.PanicOnError(g)
	w.mutex.Lock()
	previous := w.issues
	previousFavorites := w.favorites
	previousWatched := w.watched
	w.issues = issues
	w.favorites = favorites
	w.watched = watched
	w.mutex.Unlock()

	if !slices.EqualFunc(previous, issues, sameIssue) ||
		!slices.EqualFunc(previousFavorites, favorites, samePage) ||
		!slices.EqualFunc(previousWatched, watched, samePage) {
		w.notifier.Notify()
	}
}
