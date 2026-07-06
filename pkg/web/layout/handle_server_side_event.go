package layout

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"time"
)

func HandleServerSideEvent(
	n face.EventNotifier,
	render func(
		http.ResponseWriter,
		http.Flusher,
	),
) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		flusher, okay := w.(http.Flusher)

		if !okay {
			http.Error(
				w,
				"streaming not supported",
				http.StatusInternalServerError,
			)

			return
		}

		w.Header().Set(constant.ContentType, "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		subscription := n.Subscribe()
		defer n.Unsubscribe(subscription)
		render(w, flusher)

		for {
			select {
			case <-r.Context().Done():
				return
			case <-subscription:
				time.Sleep(1 * time.Second)
				drain(subscription)
				render(w, flusher)
			}
		}
	}
}
