package watcher

import "github.com/mattermost/mattermost/server/public/model"

func (w *Watcher) author(identifier string) string {
	var result *model.User
	e := w.attempt(
		func() error {
			var f error
			result, f = w.client.User(identifier)

			return f
		},
	)

	if e != nil {
		return identifier
	}

	return result.Username
}
