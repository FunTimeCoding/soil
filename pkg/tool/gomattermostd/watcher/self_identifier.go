package watcher

import "github.com/mattermost/mattermost/server/public/model"

func (w *Watcher) selfIdentifier() string {
	w.once.Do(
		func() {
			var result *model.User
			e := w.attempt(
				func() error {
					var f error
					result, f = w.client.Me()

					return f
				},
			)

			if e != nil {
				w.report(e)

				return
			}

			w.self = result.Id
		},
	)

	return w.self
}
