package navigation

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (w *Watcher) Observe(r *message.Reply) {
	v := response.NewNavigation()

	if notation.DecodeBytes(r.Parameters, v) != nil || v.Context == "" {
		return
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	switch r.Method {
	case browserConstant.LadybirdNavigationStartedEvent:
		w.started[v.Context]++
		w.open[v.Context] = true
	case browserConstant.LadybirdFragmentNavigatedEvent, browserConstant.LadybirdHistoryUpdatedEvent:
		w.started[v.Context]++
		w.open[v.Context] = false
		w.locators[v.Context] = v.Locator
	case browserConstant.LadybirdLoadEvent,
		browserConstant.LadybirdNavigationFailedEvent,
		browserConstant.LadybirdNavigationAbortedEvent:
		w.open[v.Context] = false
		w.locators[v.Context] = v.Locator
	default:
		return
	}

	close(w.changed)
	w.changed = make(chan struct{})
}
